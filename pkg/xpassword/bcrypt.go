package xpassword

import (
	"golang.org/x/crypto/bcrypt"
)

// BcryptHash hashes a plaintext password using bcrypt.
// Returns the hashed password or an error if hashing fails.
func BcryptHash(plainPass string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainPass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// BcryptCompare compares a plaintext password with a hashed password.
// Returns nil if they match, otherwise returns an error.
func BcryptCompare(plainPass, hashPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(plainPass))
}
