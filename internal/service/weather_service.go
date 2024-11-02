package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/cortezzIP/Weather-API/internal/config"
	"github.com/cortezzIP/Weather-API/internal/model"
)

type WeatherService struct {
	redisClient *redis.Client
	httpClient  *http.Client
	baseURL     string
	apiKey      string
}

func NewWeatherService(cfg *config.WeatherServiceConfig, httpClient *http.Client, redisClient *redis.Client) *WeatherService {
	return &WeatherService{
		redisClient: redisClient,
		httpClient:  httpClient,
		baseURL: cfg.APIBaseUrl,
		apiKey:  cfg.APIKey,
	}
}

func (s *WeatherService) GetCurrentWeather(location string) (*model.Weather, error) {
	ctx := context.Background()

	var weather model.Weather

	location = strings.TrimSpace(location)
	re, _ := regexp.Compile(`[^a-zA-Z]+`)
	location = re.ReplaceAllString(location, "")
	location = strings.ToLower(location)
	
	val, err := s.redisClient.Get(ctx, location).Result()
	if err == nil {
		err := json.Unmarshal([]byte(val), &weather)
		if err != nil {
			return nil, err
		}
		
		return &weather, nil
	}

	req, err := http.NewRequest("GET", s.baseURL, nil)
	if err != nil {
		slog.Error("GetCurrentWeather(): " + err.Error())
		return nil, err
	}

	req.URL.RawQuery = url.Values{
		"key": {s.apiKey},
		"q":   {location},
	}.Encode()

	resp, err := s.httpClient.Do(req)
	if err != nil {
		slog.Error("GetCurrentWeather(): " + err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 400 {
		return nil, errors.New("no matching location found")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("GetCurrentWeather(): " + err.Error())
		return nil, err
	}

	err = json.Unmarshal(body, &weather)
	if err != nil {
		slog.Error("GetCurrentWeather(): " + err.Error())
		return nil, err
	}

	err = s.redisClient.Set(ctx, location, weather, time.Minute * 10).Err()
	if err != nil {
		return nil, err
	}

	return &weather, nil
}
