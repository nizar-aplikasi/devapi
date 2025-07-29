// File: models/user.go
package models

import (
	"devapi/utils/crypto"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Fullname string    `json:"fullname"`
	NoTelp   string    `json:"notelp"`
	OrgName  string    `json:"orgname"`
	Role     string    `json:"role"`
}

func (u *User) CheckPassword(password string) bool {
	return crypto.VerifyPassword(password, u.Password)
}
