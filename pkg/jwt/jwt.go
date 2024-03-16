package jwt

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	key []byte
	exp time.Duration
}

type JWTClaims struct {
	jwt.RegisteredClaims
}

type JWTManager interface {
	GenerateToken(id string) (string, error)
	ParseToken(token string) (string, error)
}

// GenerateToken generates a JWT token.
//
// It takes the user ID as a parameter and returns the generated token string and an error.
func (j *JWT) GenerateToken(id string) (string, error) {
	now := time.Now()
	exp := now.Add(j.exp)

	data := JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ID:        id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	return token.SignedString(j.key)
}

// ParseToken parses the JWT token and returns the ID or an error.
//
// It takes a tokenStr string as a parameter and returns a string and an error.
func (j *JWT) ParseToken(tokenStr string) (string, error) {
	claims := &JWTClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return j.key, nil
	})

	if err != nil {
		return "", err
	}

	if claims.Valid() != nil {
		return "", jwt.ErrTokenInvalidClaims
	}

	return claims.ID, nil
}

// NewJWT creates a new JWT with the given key and expiration days.
//
// Parameters:
//
//	key string - the key to use for the JWT
//	days string - the expiration days as a string
//
// Returns:
//
//	*JWT - the newly created JWT
//	error - an error if the expiration days cannot be converted to an integer
func NewJWT(key, days string) (*JWT, error) {
	atoi, err := strconv.Atoi(days)
	if err != nil {
		atoi = 14
	}

	return &JWT{
		key: []byte(key),
		exp: time.Hour * 24 * time.Duration(atoi),
	}, err
}
