package crypto

import (
	"testing"
)

func TestHashAndVerifyPassword(t *testing.T) {
	password := "hidupbahagia"

	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("error hashing password: %v", err)
	}

	if !VerifyPassword(password, hashed) {
		t.Errorf("expected password verification to succeed")
	}

	if VerifyPassword("salahpassword", hashed) {
		t.Errorf("expected password verification to fail for wrong password")
	}
}
