package jwt

import (
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestGenerateToken(t *testing.T) {
	j := &JWT{
		exp: time.Hour,
		key: []byte("secret"),
	}

	id := "123"
	token, err := j.GenerateToken(id)

	if err != nil {
		t.Errorf("Error generating token: %v", err)
	}

	if token == "" {
		t.Error("Generated token is empty")
	}

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		t.Errorf("Invalid JWT token format: %s", token)
	}
}

func TestParseToken(t *testing.T) {
	j := &JWT{key: []byte("secret")}

	tokenStr := "valid_token"
	expectedID := "123"
	claims := &jwt.RegisteredClaims{ID: expectedID}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ = token.SignedString(j.key)

	id, err := j.ParseToken(tokenStr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if id != expectedID {
		t.Errorf("Expected ID %s, got %s", expectedID, id)
	}
}

func TestNewJWT(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		days        string
		expectedExp time.Duration
		expectedErr bool
	}{
		{
			name:        "Successful conversion",
			key:         "mykey",
			days:        "7",
			expectedExp: time.Hour * 24 * 7,
			expectedErr: false,
		},
		{
			name:        "Unsuccessful conversion",
			key:         "mykey",
			days:        "notanumber",
			expectedExp: time.Hour * 24 * 14, // Default value
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jwt, err := NewJWT(test.key, test.days)

			if test.expectedErr && err == nil {
				t.Errorf("Expected an error but got nil")
			}

			if !test.expectedErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if jwt.exp != test.expectedExp {
				t.Errorf("Expected expiration duration %v but got %v", test.expectedExp, jwt.exp)
			}
		})
	}
}
