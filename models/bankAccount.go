package models

// BankAccount Model to store info about the Bank Account
type BankAccount struct {
	AccountID      uint   `json:"account_id" gorm:"primary_key"`
	DocumentNumber string `json:"document_number"`
}
