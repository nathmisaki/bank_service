package models

import (
	"github.com/jinzhu/gorm"
)

// OperationType describes types of operations on Transactions
type OperationType struct {
	gorm.Model
	Description string
}
