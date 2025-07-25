// src: controllers/auth/auth_controller.go
package auth

import (
	"encoding/base64"
	"net/http"
	"strings"
	"time" // import time untuk menghitung expires_in

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	LoginForm string `json:"login_form" binding:"required"`
}

type AuthController struct {
	authService *AuthService
}

func NewAuthController(authService *AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) SignIn(ctx *gin.Context) {
	var loginRequest LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Decode login_form (base64)
	decoded, err := base64.StdEncoding.DecodeString(loginRequest.LoginForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid base64 encoding"})
		return
	}

	// Asumsikan format login_form adalah "username:password"
	credentials := string(decoded)
	parts := strings.Split(credentials, ":")
	if len(parts) != 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login format"})
		return
	}

	username := parts[0]
	password := parts[1]

	// Call AuthService to authenticate user
	accessToken, refreshToken, expire, err := c.authService.AuthenticateUser(username, password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Calculate expires_in by subtracting the current time from the expire time
	expiresIn := int(time.Until(expire).Seconds())

	// Format the response properly
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "operation success",
		"Data": gin.H{
			"access_token":  accessToken,
			"token_type":    "Bearer",
			"expires_in":    expiresIn, // add expires_in to response
			"refresh_token": refreshToken,
		},
	})
}
