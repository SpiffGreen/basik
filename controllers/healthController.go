package controllers

import (
	"github.com/gin-gonic/gin"
)

// Check handles the health check endpoint
func CheckHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "healthy",
	})
}
