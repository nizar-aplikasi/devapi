// File: models/login_request.go
package models

// LoginRequest model untuk menerima data login
type LoginRequest struct {
	LoginForm string `json:"login_form" binding:"required"`
}
