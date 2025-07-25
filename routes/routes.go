package routes

import (
	"devapi/features/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Initialize AuthController
	authService := auth.NewAuthService()
	authController := auth.NewAuthController(authService)

	// Register the auth route
	authGroup := r.Group("/api/v1/auth")
	{
		authGroup.POST("/signin", authController.SignIn)
	}
}
