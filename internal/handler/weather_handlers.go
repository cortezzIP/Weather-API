package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cortezzIP/Weather-API/internal/service"
)

type WeatherController struct {
	weatherService *service.WeatherService
}

func NewWeatherController(weatherService *service.WeatherService) *WeatherController {
	return &WeatherController{
		weatherService: weatherService,
	}
}

// GetCurrentWeather godoc
// @Summary      Get current weather
// @Description  Get current weather by location
// @Tags         weather
// @Accept       json
// @Produce      json
// @Param        location   query      string  true  "Location"
// @Success      200  {object}  model.Weather
// @Failure      400  {string}  "error"
// @Router       /weather/current [get]
func (h *WeatherController) GetCurrentWeather(c *gin.Context) {
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
