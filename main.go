package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nelsonmhjr/bank_service/models"
	"github.com/nelsonmhjr/bank_service/routes"
)

func main() {
	db := models.SetupModels(gin.Mode())
	r := routes.SetupRouter(db)
	r.Run(":8080")
}
