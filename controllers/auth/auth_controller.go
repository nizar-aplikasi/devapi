// controllers/auth/auth_controller.go
package authcontroller

import (
	"devapi/features/auth"
	"devapi/features/auth/dto"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService auth.Service
}

func NewAuthController(authService auth.Service) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) SignIn(ctx *gin.Context) {
	var req dto.LoginFormRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	credentials, err := decodeLoginForm(req.LoginForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenData, err := c.authService.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "operation failed",
			"Data": gin.H{"error": err.Error()},
		})
		return
	}

	expiresIn := int(time.Until(tokenData.ExpiresAt).Seconds())
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "operation success",
		"Success": true,
		"Data": gin.H{
			"access_token":  tokenData.AccessToken,
			"refresh_token": tokenData.RefreshToken,
			"token_type":    "Bearer",
			"expires_in":    expiresIn,
		},
	})
}

type credentials struct {
	Username string
	Password string
}

func decodeLoginForm(encoded string) (*credentials, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(string(decoded), ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Invalid login format")
	}
	return &credentials{Username: parts[0], Password: parts[1]}, nil
}