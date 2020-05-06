package main

import (
	"net/http"

	"github.com/nelsonmhjr/bank_service/controllers"

	"github.com/gin-gonic/gin"
	"github.com/nelsonmhjr/bank_service/models"
)

func main() {
	r := gin.Default()
	db := models.SetupModels()

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/accounts/:accountId", controllers.FindAccount)
	r.POST("/accounts", controllers.CreateAccounts)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
