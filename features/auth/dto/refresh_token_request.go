// File: features/auth/dto/refresh_token_request.go
package dto

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
