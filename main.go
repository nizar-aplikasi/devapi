package main

import (
	"devapi/config"
	"devapi/internal/server"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	appName = "DevAPI"
)

func main() {
	// Starting
	log.Println("🚀 Starting DevAPI...")

	// Initialize application
	if err := initializeApp(); err != nil {
		log.Printf("❌ Failed to initialize application: %v", err)
	}

	// Start server
	// Get PORT for info (optional log)
	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}

	swaggerURL := fmt.Sprintf("http://localhost:%s/swagger/index.html", port)

	// Informational log
	log.Printf("⚡ %s server running at %s\n", appName, swaggerURL)

	// Run server
	server.Run()
}

func initializeApp() error {
	// Load environment variables
	if err := loadEnv(); err != nil {
		return fmt.Errorf("environment loading failed: %w", err)
	}

	// Setup database
	if err := setupDatabase(); err != nil {
		return fmt.Errorf("database setup failed: %w", err)
	}

	return nil
}

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️ .env file not found, using system environment variables")
		// Continue without .env file as system env vars might be set
	}
	return nil
}

func setupDatabase() error {
	config.LoadConfig()
	if err := config.InitDB(); err != nil {
		return fmt.Errorf("database initialization failed: %w", err)
	}
	return nil
}
