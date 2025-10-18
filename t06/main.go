package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	warHandler := InitializeApp()

	router := gin.Default()
	router.GET("/wars", warHandler.GetWars)
	router.POST("/wars", warHandler.AddWar)
	router.GET("/wars/:name", warHandler.GetWarByName)

	router.Run("localhost:8080")
}
