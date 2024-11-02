package main

import (
	"net/http"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/cortezzIP/Weather-API/docs"
	"github.com/cortezzIP/Weather-API/internal/config"
	"github.com/cortezzIP/Weather-API/internal/database/redis"
	"github.com/cortezzIP/Weather-API/internal/handler"
	"github.com/cortezzIP/Weather-API/internal/router"
	"github.com/cortezzIP/Weather-API/internal/service"
)

// @title           Weather API
// @version         1.0
// @description     Another one weather api.

// @license.name  MIT
// @license.url   https://github.com/cortezzIP/Weather-API/blob/main/LICENSE

// @host      localhost:8080
// @BasePath  /api

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	cfg := config.MustLoad()

	redisClient, err := redis.ConnectToRedis(&cfg.RedisConfig)
	if err != nil {
		panic(err)
	}

	httpClient := http.Client{
		Timeout: time.Second * 30,
	}

	r := router.SetupRouter(handler.NewWeatherController(service.NewWeatherService(&cfg.WeatherServiceConfig, &httpClient, redisClient)))

	r.Run(cfg.ListenAddress)
}