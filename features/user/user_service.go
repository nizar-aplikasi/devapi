// File: features/user/user_service.go
package user

import (
	authfeature "devapi/features/auth"

	"devapi/models"
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	CreateUser(user models.User) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
	FindUserByID(id uuid.UUID) (*models.User, error)
	FindAllUsers() ([]models.User, error)
}

type UserService struct {
	repo UserRepository
	auth authfeature.Service
}

func NewUserService(repo UserRepository, authService authfeature.Service) Service {
	return &UserService{repo: repo, auth: authService}
}

// CreateUser : creates a new user
func (s *UserService) CreateUser(user models.User) (*models.User, error) {
	if user.Username == "" || user.Password == "" {
		return nil, errors.New("username and password are required")
	}

	existingUser, _ := s.repo.FindUserByUsername(user.Username)
	if existingUser != nil {
		return nil, errors.New("username already taken")
	}
	// ✅ Generate UUID baru
	user.ID = uuid.New()

	return s.repo.CreateUser(user)
}

// FindUserByUsername : finds a user by username
func (s *UserService) FindUserByUsername(username string) (*models.User, error) {
	return s.repo.FindUserByUsername(username)
}

func (s *UserService) FindUserByID(id uuid.UUID) (*models.User, error) {
	return s.repo.FindUserByID(id)
}

func (s *UserService) FindAllUsers() ([]models.User, error) {
	return s.repo.FindAllUsers()
}
