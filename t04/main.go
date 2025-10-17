package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type War struct {
	Name      string    `json:"war-name"`
	DateBegin time.Time `json:"begin-date"`
	Duration  int       `json:"total-day"`
}

var InMemoryWarData = []War{
	{Name: "Miryokefalon", DateBegin: time.Date(1176, time.September, 17, 0, 0, 0, 0, time.UTC), Duration: 1},
	{Name: "Yassıçemen", DateBegin: time.Date(1230, time.August, 10, 0, 0, 0, 0, time.UTC), Duration: 2},
}

func getWars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, InMemoryWarData)
}

func addWar(c *gin.Context) {
	var newWar War
	err := c.BindJSON(&newWar)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "request body is empty or invalid"})
		return
	}
	InMemoryWarData = append(InMemoryWarData, newWar)
	c.IndentedJSON(http.StatusCreated, InMemoryWarData)
}

func getWar(c *gin.Context) {
	name := c.Param("war-name")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range InMemoryWarData {
		if a.Name == name {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "war not found"})
}

func main() {
	router := gin.Default()
	router.GET("/wars", getWars)
	router.POST("/wars", addWar)
	router.GET("/wars/:war-name", getWar)

	router.Run("localhost:8080")
}
