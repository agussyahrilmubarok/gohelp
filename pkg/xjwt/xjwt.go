package xjwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// SecretKey used to sign and verify JWT tokens (replace with secure key in production)
var SecretKey = []byte("my-secret-key")

// Generate creates a JWT token with custom claims and expiration in minutes
func Generate(claims map[string]interface{}, expireMinutes int) (string, error) {
	tokenClaims := jwt.MapClaims{}

	// Add provided claims
	for k, v := range claims {
		tokenClaims[k] = v
	}

	tokenClaims["exp"] = time.Now().Add(time.Duration(expireMinutes) * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	return token.SignedString(SecretKey)
}

// Verify checks the JWT token and returns claims if valid
func Verify(tokenString string) (map[string]interface{}, error) {
	// Parse token and validate signing method
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims if token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result := make(map[string]interface{})
		for k, v := range claims {
			result[k] = v
		}
		return result, nil
	}

	return nil, errors.New("invalid token")
}
