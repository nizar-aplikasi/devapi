package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Controller struct {
	service Service
}

func NewUserController(service Service) *Controller {
	return &Controller{service: service}
}

// === Structs for consistent response ===

type UserResponseData struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	OrgName  string `json:"orgname"`
	Role     string `json:"role"`
}

type UserSuccessResponse struct {
	Success bool             `json:"Success"`
	Message string           `json:"Message"`
	Data    UserResponseData `json:"Data"`
}

type ErrorResponseData struct {
	Error string `json:"error"`
}

type UserErrorResponse struct {
	Success bool              `json:"Success"`
	Message string            `json:"Message"`
	Data    ErrorResponseData `json:"Data"`
}

// === Me Handler ===

func (c *Controller) Me(ctx *gin.Context) {
	claims, exists := ctx.Get("user_claims")
	if !exists {
		sendError(ctx, http.StatusUnauthorized, "Unauthorized", "Missing token claims")
		return
	}

	mapClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		sendError(ctx, http.StatusUnauthorized, "Unauthorized", "Invalid token claims")
		return
	}

	username, ok := mapClaims["username"].(string)
	if !ok || username == "" {
		sendError(ctx, http.StatusUnauthorized, "Unauthorized", "Username not found in token")
		return
	}

	user, err := c.service.FindUserByUsername(username)
	if err != nil {
		sendError(ctx, http.StatusNotFound, "User not found", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, UserSuccessResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data: UserResponseData{
			Username: user.Username,
			Fullname: user.Fullname,
			OrgName:  user.OrgName,
			Role:     user.Role,
		},
	})
}

// === Utility Function ===

func sendError(ctx *gin.Context, statusCode int, message string, errMsg string) {
	ctx.JSON(statusCode, UserErrorResponse{
		Success: false,
		Message: message,
		Data: ErrorResponseData{
			Error: errMsg,
		},
	})
}
