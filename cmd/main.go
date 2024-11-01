package main

import (
	"net/http"
	"time"

	"github.com/joho/godotenv"

	"github.com/cortezzIP/Weather-API/internal/config"
	"github.com/cortezzIP/Weather-API/internal/handler"
	"github.com/cortezzIP/Weather-API/internal/router"
	"github.com/cortezzIP/Weather-API/internal/service"
)

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

	r := router.SetupRouter(handler.NewWeatherHandle(service.NewWeatherService(&cfg.WeatherServiceConfig, &client)))

	r.Run(cfg.ListenAddress)
}