// features/auth/auth_service.go
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
