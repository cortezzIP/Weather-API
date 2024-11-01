package router

import (
	"github.com/cortezzIP/Weather-API/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(weatherService *handler.WeatherHandle) *gin.Engine {
	r := gin.Default()

	routerGroup := r.Group("/api")

	routerGroup.GET("/weather/current", weatherService.GetCurrentWeather)

	return r
}