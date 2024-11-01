package handler

import (
	"net/http"

	"github.com/cortezzIP/Weather-API/internal/service"
	"github.com/gin-gonic/gin"
)

type WeatherHandle struct {
	weatherService *service.WeatherService
}

func NewWeatherHandle(weatherService *service.WeatherService) *WeatherHandle {
	return &WeatherHandle{
		weatherService: weatherService,
	}
}

func (h *WeatherHandle) GetCurrentWeather(c *gin.Context) {
	location := c.Query("location")
	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "location is empty",
		})
		return
	}

	weather, err := h.weatherService.GetCurrentWeather(location)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, weather)
}
