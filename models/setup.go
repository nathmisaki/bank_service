package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // for use with gorm
)

// SetupModels Setup db connection and run AutoMigrate on models
func SetupModels() *gorm.DB {
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SSLMODE"))

	db, err := gorm.Open("postgres", connString)

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&BankAccount{}, &OperationType{}, &Transaction{})

	return db
}
