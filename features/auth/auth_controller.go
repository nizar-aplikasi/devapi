// File: features/auth/auth_controller.go
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

// === Structs for consistent response ===

type TokenResponseData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

type AuthSuccessResponse struct {
	Success bool              `json:"Success"`
	Message string            `json:"Message"`
	Data    TokenResponseData `json:"Data"`
}

type ErrorResponseData struct {
	Error string `json:"error"`
}

type AuthErrorResponse struct {
	Success bool              `json:"Success"`
	Message string            `json:"Message"`
	Data    ErrorResponseData `json:"Data"`
}

// === SignIn Handler ===

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
	ctx.JSON(http.StatusOK, AuthSuccessResponse{
		Success: true,
		Message: "Operation success",
		Data: TokenResponseData{
			AccessToken:  tokenData.AccessToken,
			RefreshToken: tokenData.RefreshToken,
			TokenType:    "Bearer",
			ExpiresIn:    expiresIn,
		},
	})
}

// === RefreshToken Handler ===

func (c *AuthController) RefreshToken(ctx *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendError(ctx, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	tokenData, err := c.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		sendError(ctx, http.StatusUnauthorized, "Token refresh failed", err)
		return
	}

	expiresIn := int(time.Until(tokenData.ExpiresAt).Seconds())
	ctx.JSON(http.StatusOK, AuthSuccessResponse{
		Success: true,
		Message: "Token refreshed successfully",
		Data: TokenResponseData{
			AccessToken:  tokenData.AccessToken,
			RefreshToken: tokenData.RefreshToken,
			TokenType:    "Bearer",
			ExpiresIn:    expiresIn,
		},
	})
}

// === Utility Functions ===

type credentials struct {
	Username string
	Password string
}

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

func sendError(ctx *gin.Context, statusCode int, message string, err error) {
	ctx.JSON(statusCode, AuthErrorResponse{
		Success: false,
		Message: message,
		Data: ErrorResponseData{
			Error: err.Error(),
		},
	})
}
