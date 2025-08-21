package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hotel struct
type Hotel struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	City     string  `json:"city"`
	Stars    int     `json:"stars"`
	PricePer float64 `json:"price_per_night"`
}

// In-memory database
var hotels = []Hotel{
	{ID: "1", Name: "Grand Palace", City: "Bangkok", Stars: 5, PricePer: 250.50},
	{ID: "2", Name: "Sea View Resort", City: "Phuket", Stars: 4, PricePer: 180.00},
}

func getHotels(c *gin.Context) {
	cityQuery := c.Query("city")

	if cityQuery != "" {
		filter := []Hotel{}
		for _, hotel := range hotels {
			if hotel.City == cityQuery {
				filter = append(filter, hotel)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, hotels)
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "healthy"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/hotels", getHotels)
	}

	r.Run(":8080")
}
