package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nelsonmhjr/bank_service/models"
)

// TransactionToCreate validates input data for CreateTransactions
type TransactionToCreate struct {
	AccountID       uint    `json:"account_id" binding:"required"`
	OperationTypeID uint    `json:"operation_type" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
}

// CreateTransactions handles POST /transactions requests
// creates a Transaction with the given data using TransactionToCreate
func CreateTransactions(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var data TransactionToCreate
	err := c.ShouldBind(&data)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Data"})
		return
	}
	transaction := models.Transaction{BankAccountID: data.AccountID,
		OperationTypeID: data.OperationTypeID,
		Amount:          data.Amount,
		EventDate:       time.Now()}

	db.Create(&transaction)

	c.JSON(http.StatusCreated, gin.H{"status": "created", "transaction": gin.H{
		"id":         transaction.ID,
		"account_id": transaction.BankAccountID,
		"amount":     transaction.Amount,
		"event_date": transaction.EventDate}})
}
