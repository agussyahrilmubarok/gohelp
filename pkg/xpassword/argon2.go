package xpassword

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/argon2"
)

// Argon2Hash hashes a plaintext password using Argon2id.
// Returns the hashed password encoded in base64 and the salt used.
func Argon2Hash(password string) (hashBase64 string, saltBase64 string, err error) {
	salt := make([]byte, 16)
	if _, err = rand.Read(salt); err != nil {
		return "", "", err
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	return base64.StdEncoding.EncodeToString(hash), base64.StdEncoding.EncodeToString(salt), nil
}

// Argon2Compare verifies a plaintext password against a base64-encoded Argon2id hash and salt.
func Argon2Compare(password, hashBase64, saltBase64 string) error {
	hash, err := base64.StdEncoding.DecodeString(hashBase64)
	if err != nil {
		return err
	}

	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		return err
	}

	newHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	if !equalBytes(hash, newHash) {
		return errors.New("password does not match")
	}

	return nil
}
