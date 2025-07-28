// File: utils/aes_encryptor.go
package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"
)

var aesKey = []byte(os.Getenv("AES_SECRET_KEY")) // 32-byte key for AES-256

// Encrypt encrypts plainText using AES-GCM and returns a base64-encoded string.
func Encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Decrypt decrypts the base64-encoded cipherText and returns the original plain text.
func Decrypt(encryptedBase64 string) (string, error) {
	cipherData, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len(cipherData) < aesGCM.NonceSize() {
		return "", errors.New("invalid cipher data")
	}

	nonce := cipherData[:aesGCM.NonceSize()]
	cipherText := cipherData[aesGCM.NonceSize():]

	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
