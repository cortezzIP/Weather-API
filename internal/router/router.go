package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"

	"github.com/cortezzIP/Weather-API/internal/handler"
)

func SetupRouter(weatherService *handler.WeatherController) *gin.Engine {
	r := gin.Default()

	routerGroup := r.Group("/api")

	routerGroup.GET("/weather/current", weatherService.GetCurrentWeather)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}