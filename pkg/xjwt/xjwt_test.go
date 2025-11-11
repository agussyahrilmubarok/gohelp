package xjwt_test

import (
	"testing"
	"time"

	"github.com/agussyahrilmubarok/gox/pkg/xjwt"
)

func TestJWTGenerateAndVerify(t *testing.T) {
	claims := map[string]interface{}{
		"user_id": 123,     // int will become float64 in MapClaims
		"role":    "admin", // string stays string
		"active":  true,    // bool
	}

	token, err := xjwt.Generate(claims, 1) // expires in 1 minute
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	if token == "" {
		t.Fatal("Generated token is empty")
	}

	verifiedClaims, err := xjwt.Verify(token)
	if err != nil {
		t.Fatalf("Failed to verify token: %v", err)
	}

	for k, v := range claims {
		verifiedValue, ok := verifiedClaims[k]
		if !ok {
			t.Errorf("Claim %q missing in verified claims", k)
			continue
		}

		switch orig := v.(type) {
		case int:
			if verifiedFloat, ok := verifiedValue.(float64); !ok || int(verifiedFloat) != orig {
				t.Errorf("Claim %q mismatch: expected %v, got %v", k, orig, verifiedValue)
			}
		default:
			if verifiedValue != orig {
				t.Errorf("Claim %q mismatch: expected %v, got %v", k, orig, verifiedValue)
			}
		}
	}

	expiredToken, _ := xjwt.Generate(claims, -1) // expired 1 minute ago
	time.Sleep(1 * time.Second)                  // ensure it's expired

	_, err = xjwt.Verify(expiredToken)
	if err == nil {
		t.Error("Expected error for expired token, got nil")
	}
}
