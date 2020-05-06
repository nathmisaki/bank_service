package models

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // for use with gorm
)

// SetupModels Setup db connection and run AutoMigrate on models
func SetupModels(ginMode string) *gorm.DB {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = os.Getenv("DB_NAME_PREFIX")
		switch ginMode {
		case gin.DebugMode:
			dbName += "_development"
		case gin.TestMode:
			dbName += "_test"
		}
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		dbName, os.Getenv("DB_PASSWORD"), os.Getenv("DB_SSLMODE"))
	fmt.Println(connStr)

	db, err := gorm.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&BankAccount{}, &OperationType{}, &Transaction{})
	var count int
	db.Model(&OperationType{}).Where("id IN (?)", []int{1, 2, 3, 4}).Count(&count)
	if count == 0 {
		var op OperationType
		op = OperationType{Description: "COMPRA A VISTA"}
		op.ID = 1
		db.Create(&op)
		op = OperationType{Description: "COMPRA PARCELADA"}
		op.ID = 2
		db.Create(&op)
		op = OperationType{Description: "SAQUE"}
		op.ID = 3
		db.Create(&op)
		op = OperationType{Description: "PAGAMENTO"}
		op.ID = 4
		db.Create(&op)
	}

	defer db.Close()

	return db
}
