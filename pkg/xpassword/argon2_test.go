package xpassword_test

import (
	"testing"

	"github.com/agussyahrilmubarok/gox/pkg/xpassword"
)

func TestArgon2HashAndCompare(t *testing.T) {
	password := "mySecretPassword123!"

	hash, salt, err := xpassword.Argon2Hash(password)
	if err != nil {
		t.Fatalf("Argon2Hash failed: %v", err)
	}

	if hash == "" || salt == "" {
		t.Fatal("Hash or salt should not be empty")
	}

	if err := xpassword.Argon2Compare(password, hash, salt); err != nil {
		t.Errorf("Argon2Compare failed for correct password: %v", err)
	}

	wrongPassword := "wrongPassword"
	if err := xpassword.Argon2Compare(wrongPassword, hash, salt); err == nil {
		t.Error("Argon2Compare should fail for incorrect password")
	}
}
