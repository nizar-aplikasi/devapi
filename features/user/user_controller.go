// File: features/user/user_controller.go
package user

import (
	"devapi/features/user/dto"
	"devapi/models"
	"devapi/utils/crypto"
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
	NoTelp   string `json:"notelp"`
	OrgName  string `json:"orgname"`
	Role     string `json:"role"`
}

type UserListSuccessResponse struct {
	Success bool               `json:"Success"`
	Message string             `json:"Message"`
	Data    []UserResponseData `json:"Data"`
}

type PaginatedUserResponse struct {
	Users      []UserResponseData `json:"users"`
	TotalCount int                `json:"total_count"`
	Page       int                `json:"page"`
	PageSize   int                `json:"page_size"`
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
// Me : mengembalikan data user yang sedang login
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
			NoTelp:   user.NoTelp,
			OrgName:  user.OrgName,
			Role:     user.Role,
		},
	})
}

// CreateUser
func (c *Controller) Create(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendError(ctx, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	hashedPassword, err := crypto.HashPassword(req.Password)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error hashing password", err.Error())
		return
	}

	newUser := models.User{
		Username: req.Username,
		Password: hashedPassword,
		Fullname: req.Fullname,
		NoTelp:   req.NoTelp,
		OrgName:  req.OrgName,
		Role:     req.Role,
	}

	createdUser, err := c.service.CreateUser(newUser)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "Failed to create user", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, UserSuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: UserResponseData{
			Username: createdUser.Username,
			Fullname: createdUser.Fullname,
			NoTelp:   createdUser.NoTelp,
			OrgName:  createdUser.OrgName,
			Role:     createdUser.Role,
		},
	})
}

// ListUsers
func (c *Controller) List(ctx *gin.Context) {
	users, err := c.service.FindAllUsers()
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Failed to fetch users", err.Error())
		return
	}

	var userResponses []UserResponseData
	for _, u := range users {
		userResponses = append(userResponses, UserResponseData{
			Username: u.Username,
			Fullname: u.Fullname,
			NoTelp:   u.NoTelp,
			OrgName:  u.OrgName,
			Role:     u.Role,
		})
	}

	ctx.JSON(http.StatusOK, UserListSuccessResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Data:    userResponses,
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
