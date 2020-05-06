package controllers

import (
	"net/http"

	"github.com/nelsonmhjr/bank_service/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AccountToCreate validates the input to CreateAccount
type AccountToCreate struct {
	DocumentNumber string `form:"document_number" json:"document_number" binding:"required"`
}

// CreateAccounts handles POST /accounts
// is used to create an BankAccount with the given data
func CreateAccounts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var data AccountToCreate
	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":         "Invalid Data",
			"error_message": err.Error(),
		})
		return
	}
	bankAccount := models.BankAccount{DocumentNumber: data.DocumentNumber}
	db.Create(&bankAccount)

	c.JSON(http.StatusCreated, gin.H{"status": "created", "account": gin.H{
		"account_id":      bankAccount.ID,
		"document_number": bankAccount.DocumentNumber}})
}

// FindAccount handles GET /accounts/:accountId
// is used to fetch information about an BankAccount
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
