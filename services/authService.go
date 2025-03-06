package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spiffgreen/basik/initializers"
	"github.com/spiffgreen/basik/models"
	"golang.org/x/crypto/bcrypt"
)

func AuthRegisterService(user *models.User) error {
	// Hash the user's password

	// Save the user to the database (implementation not shown)
	return nil
}

func AuthLoginService(email, password string, tenantID int) (string, error) {
	// Retrieve the user from the database (implementation not shown)
	var user models.User
	initializers.DB.Find(&user, "email = ? AND tenant_id = ?", email, tenantID)

	if user.Email == "" {
		return "", errors.New("invalid credentials")
	}

	// Check the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}
	// Generate a JWT token
	userID := strconv.FormatUint(uint64(user.ID), 10)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":       userID,
		"email":     user.Email,
		"role":      user.Role,
		"tenant_id": strconv.Itoa(tenantID),
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte("mysecret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
