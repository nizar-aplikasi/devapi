// File: features/auth/auth_repository.go
package auth

import (
	"database/sql"
	"devapi/config"
	"devapi/models"
	"fmt"
)

type Repository interface {
	GetUserByUsername(username string) (*models.User, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	query := `SELECT username, password, role FROM users WHERE username = $1`
	err := config.DB.QueryRow(query, username).Scan(&user.Username, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user %s not found", username)
		}
		return nil, err
	}

	return &user, nil
}
