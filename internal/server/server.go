// File: server.go	
package server

import (
	"devapi/middlewares"
	"devapi/routes"

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

	// Register all routes
	routes.RegisterRoutes(r)

	// Run server
	r.Run(":5050")
}
