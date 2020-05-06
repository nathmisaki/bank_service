package main

import (
	"github.com/nelsonmhjr/bank_service/models"
	"github.com/nelsonmhjr/bank_service/routes"
)

func main() {
	db := models.SetupModels()
	r := routes.SetupRouter(db)
	r.Run(":8080")
}
