package pg

import (
	"context"
	"fmt"
	"testing"

	"github.com/kupriyanovkk/gophkeeper/internal/server/config"
	"github.com/kupriyanovkk/gophkeeper/internal/server/model"
	"github.com/kupriyanovkk/gophkeeper/internal/server/storage/utils"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	config := config.NewConfig()
	ctx := context.Background()
	utils.RefreshTestDatabase()
	conn := utils.CreatePostgresConn(ctx, config.DatabaseDSNTest)
	defer conn.Close(ctx)

	storage := NewUserStore(conn)

	tests := []struct {
		name    string
		user    model.User
		wantErr bool
	}{
		{
			name: "test user",
			user: model.User{
				Login:    "test",
				Password: "test",
			},
			wantErr: false,
		},
		{
			name: "test user conflict",
			user: model.User{
				Login:    "test",
				Password: "test",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := storage.Create(ctx, tt.user)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NotNil(t, id)
		})
	}
}

func TestUserGet(t *testing.T) {
	config := config.NewConfig()
	ctx := context.Background()
	utils.RefreshTestDatabase()
	conn := utils.CreatePostgresConn(ctx, config.DatabaseDSNTest)
	defer conn.Close(ctx)

	storage := NewUserStore(conn)
	testCases := []struct {
		name     string
		input    model.User
		wantID   bool
		preparer func(model.User, *UserPostgresStorage)
	}{
		{
			name: "User model will have ID on successful user retrieval",
			input: model.User{
				Login:    "test",
				Password: "test",
			},
			wantID: true,
			preparer: func(user model.User, storage *UserPostgresStorage) {
				storage.Create(ctx, user)
			},
		},
		{
			name: "User model will be empty on unsuccessful user retrieval",
			input: model.User{
				Login:    "test",
				Password: "test",
			},
			wantID: false,
			preparer: func(user model.User, storage *UserPostgresStorage) {
				_, _ = conn.Exec(ctx, "DELETE FROM users WHERE login=$1", user.Login)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.preparer(tc.input, storage)

			got, err := storage.Get(ctx, tc.input)
			if !tc.wantID {
				assert.Nil(t, got.ID)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, got.ID)
		})
	}
}

func TestUserUpdate(t *testing.T) {
	config := config.NewConfig()
	ctx := context.Background()
	utils.RefreshTestDatabase()
	conn := utils.CreatePostgresConn(ctx, config.DatabaseDSNTest)
	defer conn.Close(ctx)

	store := NewUserStore(conn)
	user := model.User{
		Login:    "test",
		Password: "test",
	}
	store.Create(ctx, user)

	tests := []struct {
		name    string
		user    model.User
		wantErr bool
	}{
		{
			name: "test user",
			user: model.User{
				Login:    "test",
				Password: "new_password",
			},
			wantErr: false,
		},
		{
			name: "test user conflict",
			user: model.User{
				Login:    "test1",
				Password: "test1",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := store.Update(ctx, tt.user)
			fmt.Println(tt.name, err)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
