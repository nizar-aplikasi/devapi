// File: models/user.go
package models

import "devapi/utils/crypto"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (u *User) CheckPassword(password string) bool {
	return crypto.VerifyPassword(password, u.Password)
}
