package xpassword_test

import (
	"testing"

	"github.com/agussyahrilmubarok/gox/pkg/xpassword"
)

func TestPBKDF2HashAndCompare(t *testing.T) {
	password := "mySecretPassword123!"

	hash, salt, err := xpassword.PBKDF2Hash(password)
	if err != nil {
		t.Fatalf("PBKDF2Hash failed: %v", err)
	}

	if hash == "" || salt == "" {
		t.Fatal("Hash or salt should not be empty")
	}

	if err := xpassword.PBKDF2Compare(password, hash, salt); err != nil {
		t.Errorf("PBKDF2Compare failed for correct password: %v", err)
	}

	wrongPassword := "wrongPassword"
	if err := xpassword.PBKDF2Compare(wrongPassword, hash, salt); err == nil {
		t.Error("PBKDF2Compare should fail for incorrect password")
	}
}
