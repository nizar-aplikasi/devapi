// src: features/user/user_controller.go
package user

import (
	"net/http"
	"devapi/models"
	"github.com/gin-gonic/gin"
)

// Define the LoginRequest struct for binding the incoming JSON request
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserController struct {
	userService *UserService
}

// NewUserController membuat instance baru UserController
func NewUserController(userService *UserService) *UserController {
	return &UserController{userService: userService}
}

// Register handler untuk registrasi pengguna
func (u *UserController) Register(c *gin.Context) {
	var registerRequest RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Buat user baru melalui service
	user := models.User{
		Username: registerRequest.Username,
		Password: registerRequest.Password,
	}

	createdUser, err := u.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Kirim respons sukses
	c.JSON(http.StatusCreated, gin.H{
		"id":       createdUser.ID,
		"username": createdUser.Username,
	})
}
