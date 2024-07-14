package services

import (
	"log"
	"net/http"
	"strings"
	"time"

	jsonmodels "GitHub.com/weatherAPI/models/json"
	xmlmodels "GitHub.com/weatherAPI/models/xml"
	"GitHub.com/weatherAPI/utils"
    "GitHub.com/weatherAPI/data"
)

// Cache the weather data every 10 minutes
var cache = utils.NewCache(10 * time.Minute)

func FetchWeatherData(city string) (jsonmodels.City, int, error) {
    // Check if the weather data is already cached
    if weatherData, found := cache.Get(city); found {
        log.Printf("Data found in cache: %s", city)
        return weatherData.(jsonmodels.City), http.StatusOK, nil
    }

    log.Printf("Data not found in cache: %s", city)
    log.Printf("Fetching data from the API for city: %s", city)

    // If the weather data is not cached, fetch it from the API
    url, exists := data.CityUrls[city]
    if !exists {
        log.Println(utils.ErrCityNotFound)
        return jsonmodels.City{}, http.StatusNotFound, utils.ErrCityNotFound
    }

    xmlData, err := utils.FetchXMLData(url)
    if err != nil {
        log.Printf("Error fetching or validating XML data for city %s: %v", city, err)
        return jsonmodels.City{}, http.StatusInternalServerError, err
    }

    weatherXmlData, err := utils.ParseXML(xmlData)
    if err != nil {
        log.Printf("Error parsing XML data for city %s: %v", city, err)
        return jsonmodels.City{}, http.StatusInternalServerError, err
    }

    weatherData, err := convertToCity(weatherXmlData)
    if err != nil {
        log.Printf("Error converting XML data to City struct for city %s: %v", city, err)
        return jsonmodels.City{}, http.StatusInternalServerError, err
    }

    // Save the weather data to the cache
    cache.Set(city, weatherData)
    log.Printf("Data saved in cache for city: %s", city)

    return weatherData, http.StatusOK, nil
}

// Converts the XML data to the City struct
func convertToCity(weatherXmlData xmlmodels.Data) (jsonmodels.City, error) {
    weatherList := make([]jsonmodels.Weather, len(weatherXmlData.MetData))

    for i, metData := range weatherXmlData.MetData {
        weatherList[i] = jsonmodels.Weather{
            Day:   strings.Split(metData.Day, " ")[0],
            Valid: metData.Valid,
            Icon:  weatherXmlData.IconUrlBase + metData.Icon + "." + weatherXmlData.IconFormat,
            Temp: jsonmodels.WeatherTemp{
                Low:  metData.TempLow,
                Max:  metData.TempMax,
                Unit: metData.TempUnit,
            },
            Wind: jsonmodels.WeatherWind{
                Icon:          weatherXmlData.IconUrlBase + metData.WindIcon + "." + weatherXmlData.IconFormat,
                Direction:     metData.WindDir,
                DirectionLong: metData.WindDirLong,
                Speed:         metData.WindSpeed,
                Unit:          metData.WindUnit,
            },
        }
    }

    //Convert the date and time of the last update from String to Int64
    updatedAt, err := time.Parse("02.01.2006 15:04 MST", weatherXmlData.MetData[0].TsUpdated)
    if err != nil {
        return jsonmodels.City{}, utils.ErrConvertToCity
    }

    return jsonmodels.City{
        City:      weatherXmlData.MetData[0].DomainShortTitle,
        Country:   weatherXmlData.MetData[0].DomainCountryCode,
        UpdatedAt: updatedAt.Unix(),
        Weather:   weatherList,
    }, nil
}