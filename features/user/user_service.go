// src: features/user/user_service.go
package user

import (
	"devapi/models"
	"devapi/features/auth"
	"errors"
)

// UserService menangani logika terkait pengguna
type UserService struct {
	repo UserRepository
	auth auth.AuthService
}

// NewUserService membuat instansi baru dari UserService
func NewUserService(repo UserRepository, authService auth.AuthService) *UserService {
	return &UserService{repo: repo, auth: authService}
}

// CreateUser membuat pengguna baru
func (s *UserService) CreateUser(user models.User) (*models.User, error) {
	if user.Username == "" || user.Password == "" {
		return nil, errors.New("username and password are required")
	}

	existingUser, _ := s.repo.FindUserByUsername(user.Username)
	if existingUser != nil {
		return nil, errors.New("username already taken")
	}

	// Simulasi penyimpanan ke DB
	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
