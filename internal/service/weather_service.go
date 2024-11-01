package service

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/cortezzIP/Weather-API/internal/config"
	"github.com/cortezzIP/Weather-API/internal/model"
)

type WeatherService struct {
	client  http.Client
	baseURL string
	apiKey  string
}

func NewWeatherService(cfg *config.WeatherServiceConfig, client *http.Client) *WeatherService {
	return &WeatherService{
		client:  *client,
		baseURL: cfg.APIBaseUrl,
		apiKey:  cfg.APIKey,
	}
}

func (s *WeatherService) GetCurrentWeather(location string) (*model.Weather, error) {
	req, err := http.NewRequest("GET", s.baseURL, nil)
	if err != nil {
		slog.Error("GetCurrentWeather(): " + err.Error())
		return nil, err
	}

	req.URL.RawQuery = url.Values{
		"key": {s.apiKey},
		"q":   {location},
	}.Encode()

	resp, err := s.client.Do(req)
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

	var weather model.Weather

	err = json.Unmarshal(body, &weather)
	if err != nil {
		slog.Error("GetCurrentWeather(): " + err.Error())
		return nil, err
	}

	return &weather, nil
}
