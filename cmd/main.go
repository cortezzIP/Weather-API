package main

import (
	"net/http"
	"time"

	"github.com/joho/godotenv"

	"github.com/cortezzIP/Weather-API/internal/config"
	"github.com/cortezzIP/Weather-API/internal/handler"
	"github.com/cortezzIP/Weather-API/internal/router"
	"github.com/cortezzIP/Weather-API/internal/service"
	_ "github.com/cortezzIP/Weather-API/docs"
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

	client := http.Client{
		Timeout: time.Second * 30,
	}

	r := router.SetupRouter(handler.NewWeatherController(service.NewWeatherService(&cfg.WeatherServiceConfig, &client)))

	r.Run(cfg.ListenAddress)
}