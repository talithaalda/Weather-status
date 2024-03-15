package main

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

// Data struct to hold water and wind values
type Data struct {
	Water      int    `json:"water"`
	Wind       int    `json:"wind"`
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
}

func main() {
	// Initialize Gin
	router := gin.Default()

	// Serve static files (CSS and JS)
	router.Static("/static", "./static")

	// Load HTML templates from the templates directory
	router.LoadHTMLGlob("templates/*.html")

	// Route for serving HTML page
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	// Route for providing data JSON
	router.GET("/data", func(c *gin.Context) {
		data := fetchData()
		c.JSON(200, data)
	})

	// Run server
	router.Run(":8080")
}

func fetchData() Data {
	// Generate random values for water and wind
	water := rand.Intn(100) + 1
	wind := rand.Intn(100) + 1

	// Determine status for water and wind
	waterStatus := getStatusWater(water)
	windStatus := getStatusWWind(wind)

	// Return data
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
