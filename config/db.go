package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DB adalah variabel global untuk menyimpan koneksi database
var DB *sql.DB

// InitDB menginisialisasi koneksi ke database menggunakan konfigurasi yang ada
func InitDB() error {
	// Memastikan konfigurasi telah dimuat
	LoadConfig()

	// Membuka koneksi ke database
	var err error
	DB, err = sql.Open("postgres", AppConfig.DatabaseURL)
	if err != nil {
		log.Printf("❌ Failed to connect to database: %v", err)
		return err
	}

	// Mengecek koneksi dengan melakukan ping ke database
	if err := DB.Ping(); err != nil {
		log.Printf("❌ Database ping failed: %v", err)
		return err
	}

	fmt.Println("✅ Connected to PostgreSQL")
	return nil
}
