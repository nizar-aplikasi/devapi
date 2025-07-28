package main

import (
	"devapi/config"
	"devapi/internal/server"
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/joho/godotenv"
)

const (
	swaggerURL = "http://localhost:5050/swagger/index.html"
	appName    = "DevAPI"
)

func main() {
	// Initialize application
	if err := initializeApp(); err != nil {
		log.Fatalf("❌ Failed to initialize application: %v", err)
	}

	// Start server
	log.Printf("🚀 %s server running at %s\n", appName, swaggerURL)
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

	// Launch browser
	if err := launchBrowser(swaggerURL); err != nil {
		log.Printf("⚠️ Browser launch warning: %v", err)
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

func launchBrowser(url string) error {
	var browsers = []struct {
		os      string
		commands [][]string
	}{
		{
			"windows",
			[][]string{
				{"cmd", "/c", "start", "msedge", url},  // Microsoft Edge
				{"cmd", "/c", "start", "chrome", url},   // Google Chrome
				{"cmd", "/c", "start", url},            // Default browser
			},
		},
		{
			"darwin",
			[][]string{
				{"open", "-a", "Microsoft Edge", url},
				{"open", "-a", "Google Chrome", url},
				{"open", url},
			},
		},
		{
			"linux",
			[][]string{
				{"microsoft-edge", url},
				{"microsoft-edge-stable", url},
				{"google-chrome", url},
				{"xdg-open", url},
			},
		},
	}

	for _, browser := range browsers {
		if runtime.GOOS == browser.os {
			for _, cmdArgs := range browser.commands {
				cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
				if err := cmd.Start(); err == nil {
					return nil
				}
			}
			return fmt.Errorf("no supported browser found on %s", runtime.GOOS)
		}
	}

	return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
}