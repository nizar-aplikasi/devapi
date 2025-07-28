package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Parameter konfigurasi Argon2id
const (
	memory      = 64 * 1024 // 64 MB
	time        = 1
	threads     = 4
	keyLength   = 32
	saltLength  = 16
)

// HashPassword generates a hash using Argon2id with random salt
func HashPassword(password string) (string, error) {
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)

	// Format: $argon2id$v=19$m=65536,t=1,p=4$<salt-base64>$<hash-base64>
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s", memory, time, threads, encodedSalt, encodedHash), nil
}

// VerifyPassword checks if password matches the encoded hash
func VerifyPassword(password, encodedHash string) bool {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false
	}

	var mem uint32
	var t uint32
	var p uint8

	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &mem, &t, &p)
	if err != nil {
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false
	}

	actualHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false
	}

	compareHash := argon2.IDKey([]byte(password), salt, t, mem, p, uint32(len(actualHash)))

	return subtleCompare(compareHash, actualHash)
}

// subtleCompare does constant-time comparison
func subtleCompare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	var result byte
	for i := 0; i < len(a); i++ {
		result |= a[i] ^ b[i]
	}
	return result == 0
}
