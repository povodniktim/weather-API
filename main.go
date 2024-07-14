package main

import (
	"GitHub.com/weatherAPI/handlers"
	"GitHub.com/weatherAPI/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default()
    router.GET("/weather/:city", handlers.GetWeather)
    router.GET("/weather", handlers.GetAllWeather)
    router.Run()
}