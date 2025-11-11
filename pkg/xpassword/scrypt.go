package xpassword

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/scrypt"
)

// ScryptHash hashes a plaintext password using scrypt.
// Returns the hashed password encoded in base64 and the salt used.
func ScryptHash(password string) (hashBase64 string, saltBase64 string, err error) {
	salt := make([]byte, 16)
	if _, err = rand.Read(salt); err != nil {
		return "", "", err
	}

	hash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", "", err
	}

	return base64.StdEncoding.EncodeToString(hash), base64.StdEncoding.EncodeToString(salt), nil
}

// ScryptCompare verifies a plaintext password against a base64-encoded scrypt hash and salt.
func ScryptCompare(password, hashBase64, saltBase64 string) error {
	hash, err := base64.StdEncoding.DecodeString(hashBase64)
	if err != nil {
		return err
	}

	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		return err
	}

	newHash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return err
	}

	if !equalBytes(hash, newHash) {
		return errors.New("password does not match")
	}

	return nil
}
