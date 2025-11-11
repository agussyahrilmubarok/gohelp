package xpassword

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

// PBKDF2Hash hashes a plaintext password using PBKDF2 with SHA256.
// Returns the hashed password encoded in base64 and the salt used.
func PBKDF2Hash(password string) (hashBase64 string, saltBase64 string, err error) {
	salt := make([]byte, 16)
	if _, err = rand.Read(salt); err != nil {
		return "", "", err
	}

	hash := pbkdf2.Key([]byte(password), salt, 4096, 32, sha256.New)

	return base64.StdEncoding.EncodeToString(hash), base64.StdEncoding.EncodeToString(salt), nil
}

// PBKDF2Compare verifies a plaintext password against a base64-encoded PBKDF2 hash and salt.
func PBKDF2Compare(password, hashBase64, saltBase64 string) error {
	hash, err := base64.StdEncoding.DecodeString(hashBase64)
	if err != nil {
		return err
	}
	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		return err
	}

	newHash := pbkdf2.Key([]byte(password), salt, 4096, 32, sha256.New)

	if subtle.ConstantTimeCompare(hash, newHash) != 1 {
		return errors.New("password does not match")
	}

	return nil
}
