// src : features/auth/auth_repository.go
package auth

import (
	"devapi/config"
	"devapi/models"
	"database/sql"
	"fmt"
)

// Repository interface to define the methods for the repository
type Repository interface {
	GetUserByUsername(username string) (*models.User, error)
}

// Repository struct that implements the Repository interface
type repository struct{}

// NewRepository creates a new instance of the repository
func NewRepository() Repository {
	return &repository{}
}

// GetUserByUsername fetches a user by their username
func (r *repository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	query := "SELECT username, password FROM users WHERE username = $1"
	err := config.DB.QueryRow(query, username).Scan(&user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}
