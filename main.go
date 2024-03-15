package main

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Water      int    `json:"water"`
	Wind       int    `json:"wind"`
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
}

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	router.GET("/data", func(c *gin.Context) {
		data := fetchData()
		c.JSON(200, data)
	})

	router.Run(":8080")
}

func fetchData() Data {
	water := rand.Intn(100) + 1
	wind := rand.Intn(100) + 1

	waterStatus := getStatusWater(water)
	windStatus := getStatusWWind(wind)

	return Data{
		Water:      water,
		Wind:       wind,
		WaterStatus: waterStatus,
		WindStatus:  windStatus,
	}
}

func getStatusWater(value int) string {
	if value < 5 {
		return "Aman"
	} else if value >= 6 && value <= 8 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}
func getStatusWWind(value int) string {
	if value < 6 {
		return "Aman"
	} else if value >= 7 && value <= 15 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}
