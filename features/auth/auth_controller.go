package auth

import (
	"devapi/features/auth/dto"
	"encoding/base64"
	"errors"	
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService Service
}

func NewAuthController(authService Service) *AuthController {
	return &AuthController{authService: authService}
}

// SignIn handles POST /auth/signin
func (c *AuthController) SignIn(ctx *gin.Context) {
	var req dto.LoginFormRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendError(ctx, http.StatusBadRequest, "Invalid request format", err)
		return
	}

	credentials, err := decodeLoginForm(req.LoginForm)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "Invalid login format", err)
		return
	}

	tokenData, err := c.authService.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		sendError(ctx, http.StatusUnauthorized, "Authentication failed", err)
		return
	}

	expiresIn := int(time.Until(tokenData.ExpiresAt).Seconds())
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Operation success",
		"Success": true,
		"Data": gin.H{
			"access_token":  tokenData.AccessToken,
			"refresh_token": tokenData.RefreshToken,
			"token_type":    "Bearer",
			"expires_in":    expiresIn,
		},
	})
}

// credentials struct for internal usage
type credentials struct {
	Username string
	Password string
}

// decodeLoginForm decodes a base64 encoded login string to extract username and password
func decodeLoginForm(encoded string) (*credentials, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	parts := strings.SplitN(string(decoded), ":", 2)
	if len(parts) != 2 {
		return nil, errors.New("expected format: username:password")
	}

	return &credentials{
		Username: strings.TrimSpace(parts[0]),
		Password: strings.TrimSpace(parts[1]),
	}, nil
}

// sendError standardizes error response
func sendError(ctx *gin.Context, statusCode int, message string, err error) {
	ctx.JSON(statusCode, gin.H{
		"Success": false,
		"Message": message,
		"Data": gin.H{
			"error": err.Error(),
		},
	})
}
