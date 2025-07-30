package routes

import (
	authfeature "devapi/features/auth"
	userfeature "devapi/features/user"
	"devapi/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// ✅ Tambahkan ini agar Railway tidak error saat akses root domain
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "🚆 DevAPI is running on Railway!",
			"docs":    "/swagger/index.html",
		})
	})

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
		// POST
		auth.POST("/signin", authController.SignIn)
		auth.POST("/refresh", authController.RefreshToken)
	}

	// USER routes (requires JWT middleware)
	user := r.Group("/api/v1/user")
	user.Use(middlewares.JWTAuthMiddleware())
	{
		// POST
		user.POST("/create", userController.Create)
		// GET
		user.GET("/me", userController.Me)
		user.GET("/list", userController.List)
	}
}
