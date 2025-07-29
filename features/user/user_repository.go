// File: features/user/user_repository.go
package user

import (
	"database/sql"
	"devapi/config"
	"devapi/models"
	"errors"

	"github.com/google/uuid"
)

// UserRepository menangani operasi database untuk pengguna
type UserRepository interface {
	FindUserByUsername(username string) (*models.User, error)
	FindUserByID(id uuid.UUID) (*models.User, error)
	CreateUser(user models.User) (*models.User, error)
	FindAllUsers() ([]models.User, error)
}

// userRepository adalah implementasi dari UserRepository
type userRepository struct {
	db *sql.DB
}

// NewUserRepository : membuat instansi baru dari userRepository
func NewUserRepositoryImpl() UserRepository {
	return &userRepository{
		db: config.DB,
	}
}

// FindUserByUsername : mencari user berdasarkan username
func (r *userRepository) FindUserByUsername(username string) (*models.User, error) {
	query := `SELECT id, username, password, fullname, notelp, orgname, role FROM users WHERE username = $1 LIMIT 1`
	row := r.db.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Fullname, &user.NoTelp, &user.OrgName, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindUserByID : mencari user berdasarkan UUID
func (r *userRepository) FindUserByID(id uuid.UUID) (*models.User, error) {
	query := `SELECT id, username, password, fullname, notelp, orgname, role FROM users WHERE id = $1 LIMIT 1`
	row := r.db.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Fullname, &user.NoTelp, &user.OrgName, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// CreateUser : menyimpan user baru ke database
func (r *userRepository) CreateUser(user models.User) (*models.User, error) {
	query := `INSERT INTO users (id, username, password, fullname, notelp, orgname, role) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, user.ID, user.Username, user.Password, user.Fullname, user.NoTelp, user.OrgName, user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser : mengupdate data user di database
func (r *userRepository) UpdateUser(user models.User) (*models.User, error) {
	query := `UPDATE users SET username = $1, password = $2, fullname = $3, notelp = $4, orgname = $5, role = $6 WHERE id = $7`
	_, err := r.db.Exec(query, user.Username, user.Password, user.Fullname, user.NoTelp, user.OrgName, user.Role, user.ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser : menghapus user dari database
func (r *userRepository) DeleteUser(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

// FindAllUsers : mengembalikan daftar semua pengguna
func (r *userRepository) FindAllUsers() ([]models.User, error) {
	query := `SELECT id, username, fullname, notelp, orgname, role FROM users ORDER BY username ASC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Fullname, &user.NoTelp, &user.OrgName, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
