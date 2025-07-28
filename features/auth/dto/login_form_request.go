// File: features/auth/dto/login_form_request.go
package dto

type LoginFormRequest struct {
	LoginForm string `json:"login_form" binding:"required"`
}
