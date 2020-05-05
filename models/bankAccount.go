package models

import (
	"github.com/jinzhu/gorm"
)

// BankAccount Model to store info about the Bank Account
type BankAccount struct {
	gorm.Model
	DocumentNumber string
}
