package service

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	authMiddleware "github.com/kupriyanovkk/gophkeeper/internal/server/middleware/auth"
	"github.com/kupriyanovkk/gophkeeper/internal/server/model"
	storageMock "github.com/kupriyanovkk/gophkeeper/internal/server/storage/mocks"
	cryptMock "github.com/kupriyanovkk/gophkeeper/pkg/crypt/mocks"
	jwtMock "github.com/kupriyanovkk/gophkeeper/pkg/jwt/mocks"
	pb "github.com/kupriyanovkk/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var loc, _ = time.LoadLocation("UTC")
var now = time.Now().In(loc)

func privateTestClient(t *testing.T, ctrl *gomock.Controller, uid uuid.UUID) (pb.PrivateClient, chan<- struct{}) {
	done := make(chan struct{})
	storageMock := storageMock.NewMockPrivateStorage(ctrl)

	storageMock.EXPECT().CreatePrivateData(gomock.Any(), gomock.Any()).AnyTimes().Return(model.PrivateData{}, nil)

	storageMock.EXPECT().
		GetPrivateData(gomock.Any(), gomock.Eq(model.PrivateData{ID: 1, UserID: uid})).
		AnyTimes().
		Return(model.PrivateData{}, nil)

	storageMock.EXPECT().
		GetPrivateData(gomock.Any(), gomock.Eq(model.PrivateData{ID: 0, UserID: uid})).
		AnyTimes().
		Return(model.PrivateData{}, errors.New("test"))

	storageMock.EXPECT().
		UpdatePrivateData(gomock.Any(), gomock.Eq(model.PrivateData{ID: 1, UserID: uid, Updated: now}), gomock.Eq(true)).
		AnyTimes().
		Return(model.PrivateData{}, nil)

	storageMock.EXPECT().
		UpdatePrivateData(gomock.Any(), gomock.Eq(model.PrivateData{ID: 0, UserID: uid, Updated: now}), gomock.Eq(true)).
		AnyTimes().
		Return(model.PrivateData{}, errors.New("test"))

	storageMock.EXPECT().
		DeletePrivateData(gomock.Any(), gomock.Eq(model.PrivateData{ID: 1, UserID: uid})).
		AnyTimes().
		Return(nil)
	storageMock.EXPECT().
		DeletePrivateData(gomock.Any(), gomock.Eq(model.PrivateData{ID: 0, UserID: uid})).
		AnyTimes().
		Return(errors.New("test"))

	storageMock.EXPECT().GetPrivateDataByType(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().
		Return([]model.PrivateData{
			{ID: 1},
			{ID: 2},
		}, nil)

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

	privateService := NewPrivateService(storageMock)
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			auth.UnaryServerInterceptor(authMiddleware.NewAuthMiddleware(jwtMock, cryptMock).Auth),
		),
		grpc.StreamInterceptor(
			auth.StreamServerInterceptor(authMiddleware.NewAuthMiddleware(jwtMock, cryptMock).Auth),
		),
	)

	pb.RegisterPrivateServer(server, privateService)

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

	client := pb.NewPrivateClient(conn)

	return client, done
}

func TestRegisterPrivateService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pr := storageMock.NewMockPrivateStorage(ctrl)

	s := NewPrivateService(pr)

	server := grpc.NewServer()

	s.RegisterService(server)
}

func TestCreatePrivateData(t *testing.T) {
	uid := uuid.New()
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer token")

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	client, done := privateTestClient(t, ctl, uid)
	defer close(done)

	_, err := client.CreatePrivateData(ctx, &pb.CreatePrivateDataRequest{})
	assert.NoError(t, err)
}

func TestUpdatePrivateData(t *testing.T) {
	uid := uuid.New()
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer token")

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	client, done := privateTestClient(t, ctl, uid)
	defer close(done)

	_, err := client.UpdatePrivateData(ctx, &pb.UpdatePrivateDataRequest{Id: 1, IsForce: true, Updated: timestamppb.New(now)})
	assert.NoError(t, err)

	_, err = client.UpdatePrivateData(ctx, &pb.UpdatePrivateDataRequest{Id: 0, IsForce: true, Updated: timestamppb.New(now)})
	assert.Error(t, err)
}

func TestDeletePrivateData(t *testing.T) {
	uid := uuid.New()
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer token")

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	client, done := privateTestClient(t, ctl, uid)
	defer close(done)

	_, err := client.DeletePrivateData(ctx, &pb.DeletePrivateDataRequest{Id: 1})
	assert.NoError(t, err)

	_, err = client.DeletePrivateData(ctx, &pb.DeletePrivateDataRequest{Id: 0})
	assert.Error(t, err)
}

func TestGetPrivateData(t *testing.T) {
	uid := uuid.New()
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer token")

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	client, done := privateTestClient(t, ctl, uid)
	defer close(done)

	_, err := client.GetPrivateData(ctx, &pb.GetPrivateDataRequest{Id: 1})
	assert.NoError(t, err)

	_, err = client.GetPrivateData(ctx, &pb.GetPrivateDataRequest{Id: 0})
	assert.Error(t, err)
}

func TestGetPrivateDataByType(t *testing.T) {
	uid := uuid.New()
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer token")

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	client, done := privateTestClient(t, ctl, uid)
	defer close(done)

	res, err := client.GetPrivateDataByType(ctx, &pb.GetPrivateDataByTypeRequest{})
	assert.NoError(t, err)
	assert.Len(t, res.Data, 2)
}