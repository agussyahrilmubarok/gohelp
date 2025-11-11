package xpassword_test

import (
	"testing"

	"github.com/agussyahrilmubarok/gox/pkg/xpassword"
)

func TestScryptHashAndCompare(t *testing.T) {
	password := "mySecretPassword123!"

	hash, salt, err := xpassword.ScryptHash(password)
	if err != nil {
		t.Fatalf("ScryptHash failed: %v", err)
	}

	if hash == "" || salt == "" {
		t.Fatal("Hash or salt should not be empty")
	}

	if err := xpassword.ScryptCompare(password, hash, salt); err != nil {
		t.Errorf("ScryptCompare failed for correct password: %v", err)
	}

	wrongPassword := "wrongPassword"
	if err := xpassword.ScryptCompare(wrongPassword, hash, salt); err == nil {
		t.Error("ScryptCompare should fail for incorrect password")
	}
}
