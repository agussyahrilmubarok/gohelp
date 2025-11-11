package xpassword_test

import (
	"strings"
	"testing"

	"github.com/agussyahrilmubarok/gox/pkg/xpassword"
)

func TestBcryptHashAndComparePassword(t *testing.T) {
	plain := "MySecretPassword123!"

	hashed, err := xpassword.BcryptHash(plain)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	if hashed == "" {
		t.Fatal("Hashed password is empty")
	}

	if hashed == plain {
		t.Fatal("Hashed password should not equal the plain password")
	}

	err = xpassword.BcryptCompare(plain, hashed)
	if err != nil {
		t.Fatalf("ComparePassword failed for correct password: %v", err)
	}

	wrong := "WrongPassword123!"
	err = xpassword.BcryptCompare(wrong, hashed)
	if err == nil {
		t.Fatal("ComparePassword should fail for incorrect password")
	}

	if !strings.Contains(err.Error(), "crypto/bcrypt") {
		t.Logf("ComparePassword returned error: %v", err)
	}
}
