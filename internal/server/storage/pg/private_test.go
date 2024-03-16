package pg

import (
	"context"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kupriyanovkk/gophkeeper/internal/server/config"
	"github.com/kupriyanovkk/gophkeeper/internal/server/model"
	"github.com/kupriyanovkk/gophkeeper/internal/server/storage/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreatePrivateData(t *testing.T) {
	config := config.NewConfig()
	ctx := context.Background()
	uid := uuid.New()
	utils.RefreshTestDatabase()
	conn := utils.CreatePostgresConn(ctx, config.DatabaseDSNTest)
	defer conn.Close(ctx)

	type args struct {
		ctx     context.Context
		private model.PrivateData
	}
	tests := []struct {
		name     string
		args     args
		expected model.PrivateData
		error    assert.ErrorAssertionFunc
		preparer func()
	}{
		{
			name: "create private data",
			args: args{
				ctx: ctx,
				private: model.PrivateData{
					Title:   "test",
					Content: nil,
					Type:    1,
					UserID:  uid,
					Updated: time.Time{},
					Deleted: false,
				},
			},
			error: assert.NoError,
			preparer: func() {
				utils.RefreshTestDatabase()
				row, _ := conn.Query(
					ctx, "insert into users (id, login, password) values ($1,$2,$3)", uid, "test", "test",
				)
				defer row.Close()
			},
			expected: model.PrivateData{
				ID:      1,
				Title:   "test",
				Content: nil,
				Type:    1,
				UserID:  uid,
				Updated: time.Time{},
				Deleted: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.preparer()

			storage := NewPrivateStore(conn)
			got, err := storage.CreatePrivateData(tt.args.ctx, tt.args.private)
			tt.error(t, err, fmt.Sprintf("CreatePrivateData(%v, %v)", tt.args.ctx, tt.args.private))

			assert.Equalf(t, tt.expected, got, "CreatePrivateData(%v, %v)", tt.args.ctx, tt.args.private)
		})
	}
}

func TestUpdatePrivateData(t *testing.T) {
	config := config.NewConfig()
	ctx := context.Background()
	uid := uuid.New()
	utils.RefreshTestDatabase()
	conn := utils.CreatePostgresConn(ctx, config.DatabaseDSNTest)
	defer conn.Close(ctx)

	type args struct {
		ctx     context.Context
		private model.PrivateData
	}
	tests := []struct {
		name     string
		args     args
		wantErr  assert.ErrorAssertionFunc
		preparer func()
	}{
		{
			name: "Data can be edited",
			args: args{
				ctx: ctx,
				private: model.PrivateData{
					ID:      1,
					UserID:  uid,
					Type:    1,
					Title:   "Test new",
					Content: []byte{1, 2},
					Updated: time.Time{},
					Deleted: false,
				},
			},
			wantErr: assert.NoError,
			preparer: func() {
				utils.RefreshTestDatabase()

				row, _ := conn.Query(
					ctx, "insert into users (id, login, password) values ($1,$2,$3)", uid, "test", "test",
				)
				row.Close()

				row2, _ := conn.Query(
					ctx, "insert into private (user_id, title, content, type) values ($1,$2,$3,$4)",
					uid, "test", hex.EncodeToString([]byte{10, 20}), 1,
				)
				row2.Close()
			},
		},
		{
			name: "Error will be return when no rows is edited",
			args: args{
				ctx: ctx,
				private: model.PrivateData{
					ID:      1,
					UserID:  uid,
					Type:    1,
					Title:   "Test",
					Content: []byte{1, 2},
					Updated: time.Time{},
					Deleted: false,
				},
			},
			wantErr: assert.Error,
			preparer: func() {
				utils.RefreshTestDatabase()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.preparer()
			storage := NewPrivateStore(conn)

			got, err := storage.UpdatePrivateData(tt.args.ctx, tt.args.private, true)
			if !tt.wantErr(t, err, fmt.Sprintf("UpdatePrivateData(%v, %v)", tt.args.ctx, tt.args.private)) {
				fmt.Println(got)
				assert.Equal(t, []byte{1, 2}, got.Content)
				assert.Equal(t, "Test new", got.Title)
				assert.NotEqual(t, time.Time{}, got.Updated)
			}
		})
	}
}

func TestGetPrivateDataByType(t *testing.T) {
	config := config.NewConfig()
	ctx := context.Background()
	uid := uuid.New()
	utils.RefreshTestDatabase()
	conn := utils.CreatePostgresConn(ctx, config.DatabaseDSNTest)
	defer conn.Close(ctx)

	type args struct {
		ctx        context.Context
		privateType model.PrivateDataType
		user       model.User
	}
	tests := []struct {
		name     string
		args     args
		wantErr  assert.ErrorAssertionFunc
		wantLen  int
		preparer func()
	}{
		{
			name: "Data can be retrieved",
			args: args{
				ctx:        ctx,
				privateType: model.PrivateDataType{ID: 1},
				user:       model.User{ID: &uid},
			},
			wantErr: assert.NoError,
			wantLen: 2,
			preparer: func() {
				utils.RefreshTestDatabase()

				row, _ := conn.Query(
					ctx, "insert into users (id, login, password) values ($1,$2,$3)", uid, "test", "test",
				)
				row.Close()

				row2, _ := conn.Query(
					ctx, "insert into private (user_id, title, content, type) values ($1,$2,$3,$4)",
					uid, "test", hex.EncodeToString([]byte{10, 20}), 1,
				)
				row2.Close()

				row3, _ := conn.Query(
					ctx, "insert into private (user_id, title, content, type) values ($1,$2,$3,$4)",
					uid, "test2", hex.EncodeToString([]byte{1, 1}), 1,
				)
				row3.Close()
			},
		},
		{
			name: "Data can be empty",
			args: args{
				ctx:        ctx,
				privateType: model.PrivateDataType{ID: 1},
				user:       model.User{ID: &uid},
			},
			wantLen: 0,
			wantErr: assert.NoError,
			preparer: func() {
				utils.RefreshTestDatabase()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.preparer()

			storage := NewPrivateStore(conn)
			got, err := storage.GetPrivateDataByType(tt.args.ctx, tt.args.privateType, tt.args.user)

			tt.wantErr(t, err, fmt.Sprintf("GetPrivateDataByType(%v, %v, %v)", tt.args.ctx, tt.args.privateType, tt.args.user))

			assert.Len(t, got, tt.wantLen)
		})
	}
}
