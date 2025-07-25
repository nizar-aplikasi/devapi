// File: utils/jwtutil/jwt.go
// Description: JWT utility functions for token generation and verification
package jwtutil

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecretKey = []byte("your-secret-key")
var jwtIssuer = "pintarbisnis.id"

type CustomClaims struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	UUID     string   `json:"uuid"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(username string, roles []string) (string, error) {
	uuidVal := uuid.NewString()
	claims := CustomClaims{
		Username: username,
		Roles:    roles,
		UUID:     uuidVal,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    jwtIssuer,
			Subject:   username,
			Audience:  []string{"pintarbisnis.id"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func GenerateRefreshToken() string {
	return uuid.NewString()
}

func VerifyAccessToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
