package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/kupriyanovkk/gophkeeper/internal/server/middleware/auth"
	"github.com/kupriyanovkk/gophkeeper/internal/server/model"
	"github.com/kupriyanovkk/gophkeeper/internal/server/storage"
	"github.com/kupriyanovkk/gophkeeper/pkg/failure"
	pb "github.com/kupriyanovkk/gophkeeper/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PrivateService struct {
	pb.UnimplementedPrivateServer

	storage storage.PrivateStorage
}

// RegisterService registers the PrivateService with the given grpc.ServiceRegistrar.
//
// r: The grpc.ServiceRegistrar to register the PrivateService with.
// No return value.
func (s *PrivateService) RegisterService(r grpc.ServiceRegistrar) {
	pb.RegisterPrivateServer(r, s)
}

// CreatePrivateData creates private data.
//
// ctx context.Context, req *pb.CreatePrivateDataRequest
// *pb.CreatePrivateDataResponse, error
func (s *PrivateService) CreatePrivateData(ctx context.Context, req *pb.CreatePrivateDataRequest) (*pb.CreatePrivateDataResponse, error) {
	token := ctx.Value(auth.JwtTokenCtx{}).(string)

	private := model.PrivateData{
		UserID:  uuid.MustParse(token),
		Title:   req.Title,
		Type:    req.Type,
		Content: req.Content,
		Updated: time.Now(),
	}

	result, err := s.storage.CreatePrivateData(ctx, private)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreatePrivateDataResponse{
		Id:      result.ID,
		Title:   result.Title,
		Type:    result.Type,
		Updated: timestamppb.New(result.Updated),
		Deleted: result.Deleted,
	}, nil
}

// UpdatePrivateData updates private data.
//
// It takes a context and a request to update private data, and returns a response or an error.
func (s *PrivateService) UpdatePrivateData(ctx context.Context, req *pb.UpdatePrivateDataRequest) (*pb.UpdatePrivateDataResponse, error) {
	token := ctx.Value(auth.JwtTokenCtx{}).(string)

	private := model.PrivateData{
		UserID:  uuid.MustParse(token),
		ID:      req.Id,
		Title:   req.Title,
		Type:    req.Type,
		Content: req.Content,
		Updated: time.Now(),
	}

	_, err := s.storage.UpdatePrivateData(ctx, private)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		if errors.Is(err, failure.ErrCouldNotUpdatePrivateData) {
			return nil, status.Error(codes.FailedPrecondition, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UpdatePrivateDataResponse{}, nil
}

// GetPrivateData retrieves private data for a user.
//
// ctx context.Context, req *pb.GetPrivateDataRequest
// *pb.GetPrivateDataResponse, error
func (s *PrivateService) GetPrivateData(ctx context.Context, req *pb.GetPrivateDataRequest) (*pb.GetPrivateDataResponse, error) {
	token := ctx.Value(auth.JwtTokenCtx{}).(string)

	private := model.PrivateData{
		UserID: uuid.MustParse(token),
		ID:     req.Id,
	}

	result, err := s.storage.GetPrivateData(ctx, private)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetPrivateDataResponse{
		Id:      result.ID,
		Title:   result.Title,
		Type:    result.Type,
		Updated: timestamppb.New(result.Updated),
		Deleted: result.Deleted,
	}, nil
}

// DeletePrivateData deletes private data.
//
// ctx context.Context, req *pb.DeletePrivateDataRequest
// *pb.DeletePrivateDataResponse, error
func (s *PrivateService) DeletePrivateData(ctx context.Context, req *pb.DeletePrivateDataRequest) (*pb.DeletePrivateDataResponse, error) {
	token := ctx.Value(auth.JwtTokenCtx{}).(string)

	private := model.PrivateData{
		UserID: uuid.MustParse(token),
		ID:     req.Id,
	}

	err := s.storage.DeletePrivateData(ctx, private)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeletePrivateDataResponse{}, nil
}

// GetPrivateDataByType retrieves private data by type.
//
// ctx context.Context, req *pb.GetPrivateDataByTypeRequest
// *pb.GetPrivateDataByTypeResponse, error
func (s *PrivateService) GetPrivateDataByType(ctx context.Context, req *pb.GetPrivateDataByTypeRequest) (*pb.GetPrivateDataByTypeResponse, error) {
	token := ctx.Value(auth.JwtTokenCtx{}).(string)
	userID, errParse := uuid.Parse(token)
	if errParse != nil {
		return nil, status.Error(codes.Internal, errParse.Error())
	}

	user := model.User{
		ID: &userID,
	}
	dataType := model.PrivateDataType{
		ID: uint(req.TypeId),
	}
	result, err := s.storage.GetPrivateDataByType(ctx, dataType, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	privateData := make([]*pb.PrivateData, 0, len(result))
	for _, v := range result {
		privateData = append(privateData, &pb.PrivateData{
			Id:      v.ID,
			Content: v.Content,
			Title:   v.Title,
			Type:    v.Type,
			Updated: timestamppb.New(v.Updated),
			Deleted: v.Deleted,
		})
	}

	return &pb.GetPrivateDataByTypeResponse{
		Data: privateData,
	}, nil
}

// NewPrivateService creates a new PrivateService with the given private storage.
// It takes a storage.PrivateStorage parameter and returns a pointer to PrivateService.
func NewPrivateService(s storage.PrivateStorage) *PrivateService {
	return &PrivateService{
		storage: s,
	}
}
