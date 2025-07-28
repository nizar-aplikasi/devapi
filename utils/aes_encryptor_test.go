// File: utils/aes_encryptor_test.go
package utils

import (
	"os"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	// Simulasikan AES_SECRET_KEY environment (panjang = 32 byte)
	key := "0123456789abcdef0123456789abcdef"
	os.Setenv("AES_SECRET_KEY", key)
	aesKey = []byte(key) // reset manual agar match dengan env

	plain := "hello secret world"

	encrypted, err := Encrypt(plain)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	decrypted, err := Decrypt(encrypted)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	if decrypted != plain {
		t.Errorf("Decrypted value mismatch. got=%s, want=%s", decrypted, plain)
	}
}

func TestDecryptInvalidInput(t *testing.T) {
	os.Setenv("AES_SECRET_KEY", "0123456789abcdef0123456789abcdef")
	aesKey = []byte(os.Getenv("AES_SECRET_KEY"))

	_, err := Decrypt("not-base64")
	if err == nil {
		t.Error("Expected error for invalid base64 input")
	}
}
