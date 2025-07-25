// File: internal/server/server.go
package server

import (
	"devapi/routes"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	// Serve Swagger static
	r.Static("/swagger", "./static/swagger")

	// Register all routes
	routes.RegisterRoutes(r)

	// Run server
	r.Run(":5050")
}
