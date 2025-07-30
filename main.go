package main

import (
	"devapi/config"
	"devapi/internal/server"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/joho/godotenv"
)

const appName = "DevAPI"

func main() {
	log.Println("🚀 Starting DevAPI...")

	if err := initializeApp(); err != nil {
		log.Fatalf("❌ Failed to initialize application: %v", err)
	}

	// Build Swagger URL
	port := config.AppConfig.AppPort
	swaggerURL := fmt.Sprintf("http://localhost:%s/swagger/index.html", port)

	log.Printf("⚡ %s server running at %s\n", appName, swaggerURL)

	// Launch browser (only for development)
	if config.AppConfig.AppEnv == "development" {
		if err := launchBrowser(swaggerURL); err != nil {
			log.Printf("⚠️ Browser launch warning: %v", err)
		}
	}

	// Start HTTP server
	server.Run()
}

func initializeApp() error {
	// Load environment from .env if available
	if err := loadEnv(); err != nil {
		return fmt.Errorf("load env: %w", err)
	}

	// Initialize AppConfig from env
	config.LoadConfig()

	// Setup DB
	if err := config.InitDB(); err != nil {
		return fmt.Errorf("init DB: %w", err)
	}

	return nil
}

func loadEnv() error {
	envFile := ".env"
	if os.Getenv("APP_ENV") == "production" {
		log.Println("📦 Production mode: Using Railway env vars")
		return nil // Tidak perlu load file
	}

	log.Println("🧪 Development mode: Loading .env file")
	if err := godotenv.Load(envFile); err != nil {
		log.Println("⚠️ .env file not found, fallback to system env")
	}
	return nil
}

func launchBrowser(url string) error {
	browsers := map[string][][]string{
		"windows": {
			{"cmd", "/c", "start", "msedge", url},
			{"cmd", "/c", "start", "chrome", url},
			{"cmd", "/c", "start", url},
		},
		"darwin": {
			{"open", "-a", "Microsoft Edge", url},
			{"open", "-a", "Google Chrome", url},
			{"open", url},
		},
		"linux": {
			{"microsoft-edge", url},
			{"microsoft-edge-stable", url},
			{"google-chrome", url},
			{"xdg-open", url},
		},
	}

	if cmds, ok := browsers[runtime.GOOS]; ok {
		for _, args := range cmds {
			if err := exec.Command(args[0], args[1:]...).Start(); err == nil {
				return nil
			}
		}
		return fmt.Errorf("no supported browser found for %s", runtime.GOOS)
	}

	return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
}
