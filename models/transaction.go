package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Transaction its the main table that describes all the transactions on the Bank
type Transaction struct {
	gorm.Model
	BankAccountID   uint
	BankAccount     BankAccount
	OperationTypeID uint
	OperationType   OperationType
	Amount          float64
	EventDate       time.Time
}
