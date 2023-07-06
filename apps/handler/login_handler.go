package handler

import (
	"depogunabangunan/apps/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("secret_key") // Change this with your own secret key

// LoginHandler handles the login request
func LoginHandler(c *gin.Context) {
	// Parse the JSON request body
	var userLogin model.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate the login credentials (replace this with your own validation logic)
	if userLogin.Username != "admin" || userLogin.Pass != "admin123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userLogin.Username,
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the token as JSON response
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
