package handlers

import (
	"net/http"
    
	"GitHub.com/weatherAPI/data"
	jsonmodels "GitHub.com/weatherAPI/models/json"
	"GitHub.com/weatherAPI/responses"

	"GitHub.com/weatherAPI/services"

	"github.com/gin-gonic/gin"
)

func GetWeather(c *gin.Context) {
    city := c.Param("city")
    weatherData, status, err := services.FetchWeatherData(city)
    if err != nil {
        responses.SendJSONError(c.Writer, status, err.Error())
        return
    }

    responses.SendJSONSuccess(c.Writer, http.StatusOK, weatherData)
}

func GetAllWeather(c *gin.Context) {
	var allWeatherData []jsonmodels.City

	// Iterate over cityUrls keys
	for city := range data.CityUrls {
		weatherData, status, err := services.FetchWeatherData(city)
		if err != nil {
			responses.SendJSONError(c.Writer, status, err.Error())
			return
		}
		allWeatherData = append(allWeatherData, weatherData)
	}

	responses.SendJSONSuccess(c.Writer, http.StatusOK, allWeatherData)
}
