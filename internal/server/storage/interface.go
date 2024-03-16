package storage

import (
	"context"

	"github.com/kupriyanovkk/gophkeeper/internal/server/model"
)

type UserStorage interface {
	Create(ctx context.Context, user model.User) (model.User, error)
	Update(ctx context.Context, user model.User) (model.User, error)
	Get(ctx context.Context, user model.User) (model.User, error)
}

type PrivateStorage interface {
	CreatePrivateData(ctx context.Context, private model.PrivateData) (model.PrivateData, error)
	UpdatePrivateData(ctx context.Context, private model.PrivateData, isForce bool) (model.PrivateData, error)
	GetPrivateData(ctx context.Context, private model.PrivateData) (model.PrivateData, error)
	DeletePrivateData(ctx context.Context, private model.PrivateData) error
	GetPrivateDataByType(ctx context.Context, privateType model.PrivateDataType, user model.User) ([]model.PrivateData, error)
}
