// src: features/user/user_repository.go
package user

import (
	"devapi/models"
	"devapi/config"	
	"errors"
)

// UserRepository interface untuk interaksi database pada pengguna
type UserRepository interface {
	CreateUser(user models.User) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
}

// UserRepositoryImpl implementasi UserRepository untuk PostgreSQL
type UserRepositoryImpl struct{}

// NewUserRepositoryImpl membuat instance baru UserRepositoryImpl
func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

// CreateUser menyimpan user baru ke database
func (r *UserRepositoryImpl) CreateUser(user models.User) (*models.User, error) {
	var createdUser models.User
	err := config.DB.QueryRow("INSERT INTO users(username, password) VALUES($1, $2) RETURNING id, username", user.Username, user.Password).
		Scan(&createdUser.ID, &createdUser.Username)
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}

// FindUserByUsername mencari user berdasarkan username
func (r *UserRepositoryImpl) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
