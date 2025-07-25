// src: features/auth/auth_service.go
package auth

import (
	"devapi/config"
	"devapi/models"
	"devapi/utils/jwtutil"
	"errors"
	"time"
)

type AuthService struct {}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) AuthenticateUser(username, password string) (string, string, time.Time, error) {
	// Your authentication logic here
	var user models.User
	err := config.DB.QueryRow("SELECT username, password FROM users WHERE username = $1", username).Scan(&user.Username, &user.Password)
	if err != nil {
		return "", "", time.Time{}, errors.New("invalid username or password")
	}

	// Assume you have password validation logic (e.g., bcrypt)
	if user.Password != password {
		return "", "", time.Time{}, errors.New("invalid username or password")
	}

	// Generate JWT tokens
	accessToken, err := jwtutil.GenerateAccessToken(username, []string{"user"})
	if err != nil {
		return "", "", time.Time{}, err
	}

	// Generate a refresh token
	refreshToken := jwtutil.GenerateRefreshToken()

	// Access token expiration (e.g., 1 hour from now)
	expire := time.Now().Add(time.Hour)

	return accessToken, refreshToken, expire, nil
}
