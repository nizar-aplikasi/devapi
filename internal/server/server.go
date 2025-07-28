package server

import (
	"devapi/middlewares"
	"devapi/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	// ✅ Tambahkan CORS middleware
	r.Use(middlewares.CORSMiddleware())

	// Serve Swagger static
	r.Static("/swagger", "./static/swagger")

	// Register all routes
	routes.RegisterRoutes(r)

	// Run server
	r.Run(":5050")
}
