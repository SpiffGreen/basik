package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spiffgreen/basik/initializers"
	"github.com/spiffgreen/basik/models"
	"github.com/spiffgreen/basik/services"
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	TenantID int    `json:"tenantID" binding:"required"`
}

func AuthRegister(c *gin.Context) {
	var body RegisterBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if exists
	var existingUser *models.User
	initializers.DB.Find(&existingUser, "email = ?", body.Email)

	if existingUser != nil && existingUser.Email != "" {
		println(existingUser)
		c.JSON(400, gin.H{"error": "Email already taken"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	body.Password = string(hashedPassword)

	// Save user
	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		TenantID: body.TenantID,
		Role:     "user",
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})
}

func AuthLogin(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		TenantID int    `json:"tenantID" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := services.AuthLoginService(loginData.Email, loginData.Password, loginData.TenantID)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
