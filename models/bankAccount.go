package models

import (
	"github.com/jinzhu/gorm"
)

// AccountToCreate validates the input to CreateAccount
type AccountToCreate struct {
	DocumentNumber string `form:"document_number" json:"document_number" binding:"required"`
}

// BankAccount Model to store info about the Bank Account
type BankAccount struct {
	gorm.Model
	DocumentNumber string
}

type Bindable interface {
	ShouldBind(interface{}) error
}

func (m *BankAccount) ValidateAndCreate(db *gorm.DB, c Bindable) error {
	var data AccountToCreate
	var err error
	err = c.ShouldBind(&data)
	if err != nil {
		return err
	}
	err = db.Create(&m).Error
	if err != nil {
		return err
	}
	return nil
}
