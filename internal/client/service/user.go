package service

import (
	"github.com/kupriyanovkk/gophkeeper/internal/client/model"
	"google.golang.org/grpc/metadata"

	pb "github.com/kupriyanovkk/gophkeeper/proto"
)

type UserService struct {
	ctx    *model.GlobalContext
	client pb.UserClient
}

// NewUserService creates a new UserService instance.
//
// It takes a GlobalContext and a UserClient as parameters and returns a pointer to UserService.
func NewUserService(ctx *model.GlobalContext, client pb.UserClient) *UserService {
	return &UserService{
		ctx:    ctx,
		client: client,
	}
}

// Register registers a user.
//
// It takes a user model as a parameter and returns an error.
func (s *UserService) Register(user model.User) error {
	result, err := s.client.Register(s.ctx.Ctx, &pb.RegisterRequest{
		Login:    user.Login,
		Password: user.Password,
	})

	if err != nil {
		return err
	}

	s.ctx.Ctx = metadata.AppendToOutgoingContext(s.ctx.Ctx, "authorization", "Bearer "+result.Token)

	return nil
}

// Login is a function that handles the user login.
//
// It takes a user model as a parameter and returns an error.
func (s *UserService) Login(user model.User) error {
	result, err := s.client.Login(s.ctx.Ctx, &pb.LoginRequest{
		Login:    user.Login,
		Password: user.Password,
	})

	if err != nil {
		return err
	}

	s.ctx.Ctx = metadata.AppendToOutgoingContext(s.ctx.Ctx, "authorization", "Bearer "+result.Token)

	return nil
}
