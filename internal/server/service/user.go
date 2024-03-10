package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/kupriyanovkk/gophkeeper/internal/server/model"
	"github.com/kupriyanovkk/gophkeeper/internal/server/storage"
	"github.com/kupriyanovkk/gophkeeper/pkg/crypt"
	"github.com/kupriyanovkk/gophkeeper/pkg/failure"
	"github.com/kupriyanovkk/gophkeeper/pkg/jwt"
	pb "github.com/kupriyanovkk/gophkeeper/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServer
	storage    storage.UserStorage
	jwtManager jwt.JWTManager
	crypter    crypt.CryptAbstract
}

// RegisterService registers the User service.
//
// r grpc.ServiceRegistrar
func (s *UserService) RegisterService(r grpc.ServiceRegistrar) {
	pb.RegisterUserServer(r, s)
}

// Register registers a new user.
//
// ctx context.Context, req *pb.RegisterRequest
// *pb.RegisterResponse, error
func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := model.User{
		Login:    req.Login,
		Password: req.Password,
	}

	result, err := s.storage.Create(ctx, user)
	if err != nil {
		if errors.Is(err, failure.ErrConflict) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	token, err := s.jwtManager.GenerateToken(strconv.FormatUint(uint64(result.ID), 10))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.RegisterResponse{Token: s.crypter.Encode(token)}, nil
}

// Login is a function to authenticate a user and generate a token.
//
// It takes a context.Context and a *pb.LoginRequest as parameters and returns a *pb.LoginResponse and an error.
func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	result, err := s.storage.Get(ctx, model.User{
		Login:    req.Login,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	token, err := s.jwtManager.GenerateToken(strconv.FormatUint(uint64(result.ID), 10))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.LoginResponse{Token: s.crypter.Encode(token)}, nil
}

// NewUserService returns a new instance of UserService.
//
// Parameters:
//
//	s - storage.UserStorage
//	m - jwt.JWTManager
//	c - crypt.CryptAbstract
//
// Return:
//
//	*UserService
func NewUserService(s storage.UserStorage, m jwt.JWTManager, c crypt.CryptAbstract) *UserService {
	return &UserService{
		storage:    s,
		jwtManager: m,
		crypter:    c,
	}
}
