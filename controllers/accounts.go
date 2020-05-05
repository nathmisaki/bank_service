package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAccounts is used to create an Account with the given data
func CreateAccounts(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}
