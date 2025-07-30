package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort     string
	AppEnv      string
	SwaggerPath string
	DatabaseURL string
}

var AppConfig Config

func LoadConfig() {
	env := getEnv("APP_ENV", "development")
	envFile := ".env.local"

	if env == "production" {
		log.Println("📦 Production mode: Using Railway env vars")
	} else {
		log.Printf("🧪 Development mode: Loading %s", envFile)
		if err := godotenv.Load(envFile); err != nil {
			log.Printf("⚠️ Failed to load %s: %v", envFile, err)
		}
	}

	AppConfig = Config{
		AppPort:     getEnv("PORT", "5050"),
		AppEnv:      env,
		SwaggerPath: getEnv("SWAGGER_PATH", "./static/swagger"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:123456@localhost:5432/db_devapi?sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
