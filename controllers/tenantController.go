package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spiffgreen/basik/initializers"
	"github.com/spiffgreen/basik/models"
)

func CreateTenant(c *gin.Context) {
	var body struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save user
	newTenant := models.Tenant{
		Name: body.Name,
	}

	if err := initializers.DB.Create(&newTenant).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create tenant: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created Tenant"})
}

func GetTenants(c *gin.Context) {
	var tenants []models.Tenant

	userRole, _ := c.Get("userRole")
	role := userRole.(string)

	if role != "admin" {
		c.JSON(500, gin.H{"error": "Not allowed to view this resource"})
		return
	}

	initializers.DB.Find(&tenants)

	c.JSON(http.StatusCreated, gin.H{"message": "Tenants", "data": tenants})
}
