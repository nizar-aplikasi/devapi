// File: features/auth/auth_service.go
package auth

import (
	"devapi/utils"
	"devapi/utils/jwtutil"
	"fmt"
	"os"
	"time"
)

type Service interface {
	Authenticate(username, password string) (*TokenData, error)
	RefreshToken(refreshToken string) (*TokenData, error)
}

type AuthService struct {
	repo Repository
}

func NewAuthService(repo Repository) Service {
	return &AuthService{repo: repo}
}

type TokenData struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

// === Authenticate ===
func (s *AuthService) Authenticate(username, password string) (*TokenData, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("user %s not found", username)
	}

	ok, err := utils.VerifyPassword(password, user.Password)
	if err != nil || !ok {
		if os.Getenv("APP_ENV") == "development" {
			return nil, fmt.Errorf("invalid password for user=%s", username)
		}
		return nil, fmt.Errorf("invalid username or password")
	}

	token, err := jwtutil.GenerateAccessToken(user.Username, []string{user.Role})
	if err != nil {
		return nil, err
	}

	refresh := jwtutil.GenerateRefreshToken()
	expiry := time.Now().Add(time.Hour)
	return &TokenData{AccessToken: token, RefreshToken: refresh, ExpiresAt: expiry}, nil
}

// === RefreshToken ===
func (s *AuthService) RefreshToken(refreshToken string) (*TokenData, error) {
	// TODO: Simpan dan validasi refresh token di database untuk keamanan lebih baik
	if refreshToken == "" {
		return nil, fmt.Errorf("refresh token is required")
	}

	// Demo saja: tidak ada validasi token, langsung generate token baru
	// Idealnya: cocokkan dengan yang disimpan, cek expired, revoked, dll

	// Contoh demo pakai default user
	user, err := s.repo.GetUserByUsername("admin")
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token (user not found)")
	}

	newAccessToken, err := jwtutil.GenerateAccessToken(user.Username, []string{user.Role})
	if err != nil {
		return nil, err
	}

	newRefreshToken := jwtutil.GenerateRefreshToken()
	expiry := time.Now().Add(time.Hour)

	return &TokenData{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		ExpiresAt:    expiry,
	}, nil
}
