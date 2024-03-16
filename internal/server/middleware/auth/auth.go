package auth

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/kupriyanovkk/gophkeeper/pkg/crypt"
	"github.com/kupriyanovkk/gophkeeper/pkg/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Auth(ctx context.Context) (context.Context, error)
}

type AuthMiddleware struct {
	jwtManager     jwt.JWTManager
	crypter        crypt.CryptAbstract
	skippedMethods []string
}

type JwtTokenCtx struct{}

// Auth is a function of AuthMiddleware that handles authentication in the context.
//
// It takes a context as a parameter and returns a context and an error.
func (m *AuthMiddleware) Auth(ctx context.Context) (context.Context, error) {
	if m.isSkippedMethod(ctx) {
		return ctx, nil
	}

	encToken, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	token, decodeErr := m.crypter.Decode(encToken)
	if decodeErr != nil {
		return nil, status.Errorf(codes.Unauthenticated, "crypter decode error: %v", decodeErr)
	}

	decodedToken, jwtDecodeErr := m.jwtManager.ParseToken(token)
	if jwtDecodeErr != nil {
		return nil, status.Errorf(codes.Unauthenticated, "jwt decode error: %v", jwtDecodeErr)
	}

	ctx = context.WithValue(ctx, JwtTokenCtx{}, decodedToken)

	return ctx, nil
}

// isSkippedMethod checks if the method is skipped.
//
// ctx context.Context
// bool
func (m *AuthMiddleware) isSkippedMethod(ctx context.Context) bool {
	method, _ := grpc.Method(ctx)
	for _, m := range m.skippedMethods {
		if m == method {
			return true
		}
	}
	return false
}

// NewAuthMiddleware creates a new AuthMiddleware instance.
//
// Parameters:
//
//	j jwt.JWTManager - the JWT manager
//	c crypt.Crypter - the crypter
//
// Return type: *AuthMiddleware
func NewAuthMiddleware(j jwt.JWTManager, c crypt.CryptAbstract) *AuthMiddleware {
	return &AuthMiddleware{
		jwtManager:     j,
		crypter:        c,
		skippedMethods: []string{"/proto.User/Register", "/proto.User/Login"},
	}
}
