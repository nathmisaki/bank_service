package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nelsonmhjr/bank_service/controllers"
)

// SetupRouter defines routes, middlewares and db variable to controllers
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	if gin.Mode() != gin.TestMode || os.Getenv("DEBUG_TEST") == "true" {
		r.Use(gin.Logger())
	}
	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Ping test
	r.GET("/ping", controllers.Ping)

	r.GET("/accounts/:accountId", controllers.FindAccount)
	r.POST("/accounts", controllers.CreateAccounts)

	r.POST("/transactions", controllers.CreateTransactions)

	return r
}
