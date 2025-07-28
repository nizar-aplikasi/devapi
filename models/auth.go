//src/models/auth.go
package models

// LoginRequest mendefinisikan struktur data untuk login
type LoginRequest struct {
	LoginForm string `json:"login_form" binding:"required"`
}
