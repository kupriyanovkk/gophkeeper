package service

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	authMiddleware "github.com/kupriyanovkk/gophkeeper/internal/server/middleware/auth"
	"github.com/kupriyanovkk/gophkeeper/internal/server/model"
	storageMock "github.com/kupriyanovkk/gophkeeper/internal/server/storage/mocks"
	cryptMock "github.com/kupriyanovkk/gophkeeper/pkg/crypt/mocks"
	"github.com/kupriyanovkk/gophkeeper/pkg/failure"
	jwtMock "github.com/kupriyanovkk/gophkeeper/pkg/jwt/mocks"
	pb "github.com/kupriyanovkk/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func userTestClient(t *testing.T, ctrl *gomock.Controller, uid uuid.UUID) (pb.UserClient, chan<- struct{}) {
	done := make(chan struct{})
	storageMock := storageMock.NewMockUserStorage(ctrl)

	storageMock.
		EXPECT().
		Get(gomock.Any(), gomock.Eq(model.User{Login: "loginOk", Password: "pass"})).
		AnyTimes().
		Return(model.User{ID: &uid, Login: "test", Password: "pass"}, nil)

	storageMock.
		EXPECT().
		Get(gomock.Any(), gomock.Eq(model.User{Login: "logineErr", Password: "pass"})).
		AnyTimes().
		Return(model.User{Login: "test", Password: "pass"}, errors.New("user login error"))

	storageMock.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(model.User{Login: "RegConflict", Password: "pass"})).
		AnyTimes().
		Return(model.User{Login: "test", Password: "pass"}, failure.ErrConflict)

	storageMock.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(model.User{Login: "RegOk", Password: "pass"})).
		AnyTimes().
		Return(model.User{ID: &uid, Login: "test1", Password: "pass"}, nil)

	storageMock.
		EXPECT().
		Update(
			gomock.Any(),
			gomock.Any(),
		).
		AnyTimes().
		Return(model.User{}, nil)

	jwtMock := jwtMock.NewMockJWTManager(ctrl)
	jwtMock.EXPECT().GenerateToken(uid.String()).AnyTimes().Return("token", nil)
	jwtMock.EXPECT().ParseToken(gomock.Any()).AnyTimes().Return(uid.String(), nil)

	cryptMock := cryptMock.NewMockCryptAbstract(ctrl)
	cryptMock.EXPECT().Encode(gomock.Any()).AnyTimes().Return("token")
	cryptMock.EXPECT().Decode(gomock.Any()).AnyTimes().Return(uid.String(), nil)

	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatal(err)
	}

	userService := NewUserService(storageMock, jwtMock, cryptMock)
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			auth.UnaryServerInterceptor(authMiddleware.NewAuthMiddleware(jwtMock, cryptMock).Auth),
		),
		grpc.StreamInterceptor(
			auth.StreamServerInterceptor(authMiddleware.NewAuthMiddleware(jwtMock, cryptMock).Auth),
		),
	)

	pb.RegisterUserServer(server, userService)

	go func() {
		if err = server.Serve(listener); err != nil && err != grpc.ErrServerStopped {
			panic(err)
		}
	}()

	conn, err := grpc.Dial(listener.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		<-done
		server.GracefulStop()
		conn.Close()
	}()

	client := pb.NewUserClient(conn)

	return client, done
}

func TestRegisterUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	us := storageMock.NewMockUserStorage(ctrl)
	j := jwtMock.NewMockJWTManager(ctrl)
	c := cryptMock.NewMockCryptAbstract(ctrl)

	s := NewUserService(us, j, c)

	server := grpc.NewServer()

	s.RegisterService(server)
}

func TestRegister(t *testing.T) {
	uid := uuid.New()
	ctx := context.WithValue(context.Background(), authMiddleware.JwtTokenCtx{}, uid)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client, done := userTestClient(t, ctrl, uid)
	defer close(done)

	_, err := client.Register(ctx, &pb.RegisterRequest{
		Login:    "RegConflict",
		Password: "pass",
	})
	assert.Error(t, err)

	_, err = client.Register(ctx, &pb.RegisterRequest{
		Login:    "RegOk",
		Password: "pass",
	})
	assert.NoError(t, err)
}

func TestLogin(t *testing.T) {
	uid := uuid.New()
	ctx := context.WithValue(context.Background(), authMiddleware.JwtTokenCtx{}, uid)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client, done := userTestClient(t, ctrl, uid)
	defer close(done)

	_, err := client.Login(ctx, &pb.LoginRequest{
		Login:    "loginOk",
		Password: "pass",
	})
	assert.NoError(t, err)

	_, err = client.Login(ctx, &pb.LoginRequest{
		Login:    "logineErr",
		Password: "pass",
	})
	assert.Error(t, err)
}
