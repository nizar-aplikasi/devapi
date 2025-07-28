package routes

import (
	authfeature "devapi/features/auth"
	userfeature "devapi/features/user"
	"devapi/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Init repositories
	authRepo := authfeature.NewRepository()
	userRepo := userfeature.NewUserRepositoryImpl() // ✅ Sudah konsisten

	// Init services
	authService := authfeature.NewAuthService(authRepo)
	userService := userfeature.NewUserService(userRepo, authService) // ✅ Cocok dengan interface

	// Init controllers
	authController := authfeature.NewAuthController(authService)
	userController := userfeature.NewUserController(userService)

	// AUTH routes (no auth required)
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/signin", authController.SignIn)
		auth.POST("/refresh", authController.RefreshToken)
	}

	// USER routes (requires JWT middleware)
	user := r.Group("/api/v1/user")
	user.Use(middlewares.JWTAuthMiddleware())
	{
		user.GET("/me", userController.Me)
	}
}
