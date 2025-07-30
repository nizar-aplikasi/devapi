package main

import (
	"devapi/config"
	"devapi/internal/server"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

const appName = "DevAPI"

func main() {
	log.Println("🚀 Starting DevAPI...")

	// Load env & app config
	config.LoadConfig()

	// Init DB
	if err := config.InitDB(); err != nil {
		log.Fatalf("❌ Failed to initialize database: %v", err)
	}

	// Build Swagger URL
	port := config.AppConfig.AppPort
	swaggerURL := fmt.Sprintf("http://localhost:%s/swagger/index.html", port)

	log.Printf("⚡ %s server running at %s\n", appName, swaggerURL)

	// Open browser only in development
	if config.AppConfig.AppEnv == "development" {
		if err := launchBrowser(swaggerURL); err != nil {
			log.Printf("⚠️ Browser launch warning: %v", err)
		}
	}

	server.Run()
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
