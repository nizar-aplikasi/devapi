// File: server.go
package server

import (
	"devapi/middlewares"
	"devapi/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// ✅ Tambahkan CORS middleware
	r.Use(middlewares.CORSMiddleware())
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.1"})

	// Serve Swagger static
	r.Static("/swagger", "./static/swagger")

	// ✅ Redirect "/" ke Swagger
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})

	// Register all routes
	routes.RegisterRoutes(r)

	// ✅ Baca dari environment PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "5050" // fallback default
	}

	// ✅ Run server dengan port dynamic
	r.Run(fmt.Sprintf(":%s", port))
}
