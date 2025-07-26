// File: main.go
package main

import (
	"devapi/config"
	"devapi/internal/server"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {
	const swaggerURL = "http://localhost:5050/swagger/index.html"

	// Setup Database connection
	if err := setupDatabase(); err != nil {
		log.Fatalf("❌ Error setting up database: %v", err)
	}

	// Inform the user about the server
	fmt.Printf("🚀 API server running at %s\n", swaggerURL)

	// Open the Swagger UI in the browser (if possible)
	if err := openBrowser(swaggerURL); err != nil {
		log.Printf("⚠️ Failed to open browser: %v", err)
	}

	// Start the server
	server.Run()
}

// setupDatabase handles the database initialization from config
func setupDatabase() error {
	// Initialize database connection from config
	// Memuat konfigurasi dari file config
	config.LoadConfig()

	// Setup koneksi database
	if err := config.InitDB(); err != nil {
		return fmt.Errorf("failed to initialize database: %v", err)
	}

	return nil
}

// openBrowser tries to open the Swagger UI in Google Chrome based on the OS.
// Returns error if it fails to open the browser
func openBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "chrome", url)
	case "darwin":
		cmd = exec.Command("open", "-a", "Google Chrome", url)
	case "linux":
		cmd = exec.Command("google-chrome", url)
	default:
		return fmt.Errorf("unsupported OS: unable to open the browser automatically")
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to open browser: %v", err)
	}
	return nil
}
