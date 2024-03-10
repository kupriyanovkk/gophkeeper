package pg

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kupriyanovkk/gophkeeper/internal/server/model"
	"github.com/kupriyanovkk/gophkeeper/pkg/failure"
)

type PrivatePostgresStorage struct {
	conn *pgx.Conn
}

// CreatePrivateData creates private data in the PrivatePostgresStorage.
//
// ctx context.Context, private model.PrivateData
// model.PrivateData, error
func (s *PrivatePostgresStorage) CreatePrivateData(ctx context.Context, private model.PrivateData) (model.PrivateData, error) {
	ctxWithTimeOut, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	sql := `
		insert into private (user_id, title, type, content, updated, deleted)
		values ($1,$2,$3,$4,$5,$6)
		returning id
	`
	err := s.conn.QueryRow(ctxWithTimeOut, sql, private.UserID, private.Title, private.Type,
		hex.EncodeToString(private.Content), private.Updated, private.Deleted,
	).Scan(&private.ID)

	if err != nil {
		return private, fmt.Errorf("failed to create private data: %w", err)
	}

	return private, nil
}

// UpdatePrivateData updates private data in the PrivatePostgresStorage.
//
// ctx - the context for the operation
// private - the private data to be updated
// (model.PrivateData, error) - returns the updated private data and an error, if any
func (s *PrivatePostgresStorage) UpdatePrivateData(ctx context.Context, private model.PrivateData) (model.PrivateData, error) {
	ctxWithTimeOut, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	p, _ := s.GetPrivateData(ctx, private)

	if p.Updated.Unix() != private.Updated.Unix() {
		return private, failure.ErrCouldNotUpdatePrivateData
	}

	sql := `
		update private
		set content = $1, updated = $2
		where user_id = $3 and title = $4 and type = $5
		returning id
	`
	err := s.conn.QueryRow(ctxWithTimeOut, sql, hex.EncodeToString(private.Content), private.Updated, private.UserID, private.Title, private.Type).Scan(&private.ID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return private, fmt.Errorf("failed to update private data: %w", err)
		}

		return private, err
	}

	return private, nil
}

// GetPrivateData retrieves private data from the PrivatePostgresStorage.
//
// ctx - the context for the operation
// private - the private data to retrieve
// model.PrivateData, error - returns the retrieved private data and any error encountered
func (s *PrivatePostgresStorage) GetPrivateData(ctx context.Context, private model.PrivateData) (model.PrivateData, error) {
	ctxWithTimeOut, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	sql := `
		select id, user_id, title, type, content, updated, deleted
		from private
		where id = $1 and user_id = $2
	`
	err := s.conn.QueryRow(ctxWithTimeOut, sql, private.ID, private.UserID).Scan(&private.ID, &private.UserID, &private.Title, &private.Type, &private.Content, &private.Updated, &private.Deleted)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return private, fmt.Errorf("failed to get private data: %w", err)
		}

		return private, err
	}

	decoded, errDecode := hex.DecodeString(string(private.Content))
	if errDecode != nil {
		return private, errDecode
	}

	private.Content = decoded

	return private, nil
}

// DeletePrivateData deletes private data from the PrivatePostgresStorage.
//
// ctx context.Context, private model.PrivateData
// error
func (s *PrivatePostgresStorage) DeletePrivateData(ctx context.Context, private model.PrivateData) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	var id int
	sql := `
		delete from private
		where id = $1 and user_id = $2
		returning id
	`
	err := s.conn.QueryRow(ctxWithTimeout, sql, private.ID, private.UserID).Scan(&id)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("failed to delete private data: %w", err)
		}

		return err
	}

	return nil
}

// GetPrivateDataByType retrieves private data of a specific type for a given user.
//
// ctx context.Context, privateType model.PrivateDataType, user model.User
// []model.PrivateData, error
func (s *PrivatePostgresStorage) GetPrivateDataByType(ctx context.Context, privateType model.PrivateDataType, user model.User) ([]model.PrivateData, error) {
	ctxWithTimeOut, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	var privateData []model.PrivateData
	sql := `
		select id, user_id, title, type, content, updated, deleted
		from private
		where user_id = $1 and type = $2
	`
	rows, err := s.conn.Query(ctxWithTimeOut, sql, user.ID, privateType.ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		private := model.PrivateData{}
		if err = rows.Scan(&private.ID, &private.UserID, &private.Title, &private.Type, &private.Content, &private.Updated, &private.Deleted); err != nil {
			return nil, err
		}

		private.Content, err = hex.DecodeString(string(private.Content))
		if err != nil {
			return nil, err
		}

		privateData = append(privateData, private)
	}

	return privateData, nil
}

// bootstrap is a function that handles the initialization of the private postgres storage.
//
// It takes a context as a parameter and returns an error.
func (s *PrivatePostgresStorage) bootstrap(ctx context.Context) error {
	tx, err := s.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
		create table if not exists private (
			id bigserial primary key,
			user_id int not null,
			title text not null,
			type int not null,
			content bytea not null,
			updated timestamp default now(),
			deleted boolean not null
		)
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `
		create index if not exists private_id_idx on private (id)
	`)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// NewPrivateStore creates a new instance of PrivatePostgresStorage and initializes it with the provided pgx.Conn.
//
// c: a pointer to the pgx.Conn object to be used for database operations.
// Returns a pointer to the newly created PrivatePostgresStorage object.
func NewPrivateStore(c *pgx.Conn) *PrivatePostgresStorage {
	storage := &PrivatePostgresStorage{
		conn: c,
	}

	err := storage.bootstrap(context.Background())
	if err != nil {
		panic(err)
	}

	return storage
}
