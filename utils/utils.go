package utils

import (
	"encoding/base64"
	"errors"
)

// DecodeBase64Login decodes a base64 encoded "username:password" string
func DecodeBase64Login(encoded string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", errors.New("failed to decode base64 login_form")
	}
	return string(decodedBytes), nil
}
