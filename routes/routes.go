// routes/routes.go
package routes

import (
	authcontroller "devapi/controllers/auth"
	authfeature "devapi/features/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	authRepo := authfeature.NewRepository()
	authService := authfeature.NewAuthService(authRepo)
	authController := authcontroller.NewAuthController(authService)

	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/signin", authController.SignIn)
	}
}
