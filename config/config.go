package config

import (
	"os"
)

// Struct untuk menyimpan konfigurasi aplikasi
type Config struct {
	AppPort     string
	AppEnv      string
	SwaggerPath string
	DatabaseURL string
}

// Global instance untuk konfigurasi
var AppConfig Config

// LoadConfig membaca variabel lingkungan (environment variables) dan menginisialisasi konfigurasi aplikasi
func LoadConfig() {
	AppConfig = Config{
		AppPort:     getEnv("APP_PORT", ":5050"),
		AppEnv:      getEnv("APP_ENV", "development"),
		SwaggerPath: getEnv("SWAGGER_PATH", "./static/swagger"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:123456@localhost:5432/db_devapi?sslmode=disable"),
	}
}

// getEnv mencoba mengambil nilai dari environment variable, jika tidak ada menggunakan nilai fallback
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
