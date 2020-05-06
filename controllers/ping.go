package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping handles GET /ping
// famous request for pong
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
