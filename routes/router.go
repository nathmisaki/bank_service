package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nelsonmhjr/bank_service/controllers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
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

	return r
}
