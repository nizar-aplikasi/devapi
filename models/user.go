package models

import (
	"database/sql"
	"errors"
	"devapi/config"
	"devapi/utils"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	// PasswordHash tidak dikirim ke response
	PasswordHash string `json:"-"`
}

// AuthenticateUser checks username and password against the database.
func AuthenticateUser(username, password string) (*User, error) {
	user := User{}

	query := `
		SELECT id, username, full_name, password_hash
		FROM users
		WHERE LOWER(username) = LOWER($1)
		LIMIT 1
	`

	err := config.DB.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.FullName,
		&user.PasswordHash,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid password")
	}

	// Clean up hash before return (optional)
	user.PasswordHash = ""

	return &user, nil
}
