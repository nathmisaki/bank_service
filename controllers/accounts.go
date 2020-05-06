package controllers

import (
	"fmt"
	"net/http"

	"github.com/nelsonmhjr/bank_service/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type accountToCreate struct {
	DocumentNumber string `form:"document_number" json:"document_number" binding:"required"`
}

// POST /accounts
// CreateAccounts is used to create an BankAccount with the given data
func CreateAccounts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var data accountToCreate
	err := c.ShouldBind(&data)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid Data"})
		return
	}
	bankAccount := models.BankAccount{DocumentNumber: data.DocumentNumber}
	db.Create(&bankAccount)

	c.JSON(http.StatusCreated, gin.H{"status": "created", "account": gin.H{
		"account_id":      bankAccount.ID,
		"document_number": bankAccount.DocumentNumber}})
}

// GET /accounts/:accountId
// FindAccount is used to fetch information about an BankAccount
func FindAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var bankAccount models.BankAccount
	err := db.Where("id = ?", c.Param("accountId")).First(&bankAccount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id":      bankAccount.ID,
		"document_number": bankAccount.DocumentNumber})
}
