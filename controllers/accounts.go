package controllers

import (
	"net/http"

	"github.com/nelsonmhjr/bank_service/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// POST /accounts
// CreateAccounts is used to create an BankAccount with the given data
func CreateAccounts(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

// GET /accounts/:accountId
// FindAccount is used to fetch information about an BankAccount
func FindAccount(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var bankAccount models.BankAccount
	if err := db.Where("id = ?", c.Param("accountId")).First(&bankAccount).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"account_id":      bankAccount.ID,
		"document_number": bankAccount.DocumentNumber})
}
