// routes/routes.go
package routes

import (
	
	authfeature "devapi/features/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	authRepo := authfeature.NewRepository()
	authService := authfeature.NewAuthService(authRepo)
	authController := authfeature.NewAuthController(authService)

	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/signin", authController.SignIn)
	}
}
