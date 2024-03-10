package pg

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/kupriyanovkk/gophkeeper/internal/server/model"
	"github.com/kupriyanovkk/gophkeeper/pkg/failure"
)

type UserPostgresStorage struct {
	conn *pgx.Conn
}

// Create creates a new user in the UserPostgresStorage.
//
// It takes a context and a user model as parameters, and returns a user model and an error.
func (s *UserPostgresStorage) Create(ctx context.Context, user model.User) (model.User, error) {
	ctxWithTimeOut, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	sql := `
		insert into users (login, password)
		values ($1, crypt($2, gen_salt('bf')))
		returning id
	`
	err := s.conn.QueryRow(ctxWithTimeOut, sql, user.Login, user.Password).Scan(&user.ID)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
				return user, failure.ErrConflict
			}
		}

		return user, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// Get retrieves a user from the UserPostgresStorage.
//
// It takes a context and a user model as parameters and returns a user model and an error.
func (s *UserPostgresStorage) Get(ctx context.Context, user model.User) (model.User, error) {
	ctxWithTimeOut, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	sql := `
		select id
		from users
		where login = $1 and password = crypt($2, password)
	`
	err := s.conn.QueryRow(ctxWithTimeOut, sql, user.Login, user.Password).Scan(&user.ID)
	if err != nil {
		return user, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *UserPostgresStorage) Update(ctx context.Context, user model.User) (model.User, error) {
	ctxWithTimeOut, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	sql := `
		update users
		set password = crypt($2, gen_salt('bf'))
		where login = $1
	`
	_, err := s.conn.Exec(ctxWithTimeOut, sql, user.Login, user.Password)
	if err != nil {
		return user, fmt.Errorf("failed to update user: %w", err)
	}
	return user, nil
}

// bootstrap bootstraps the UserPostgresStorage.
//
// ctx context.Context
// error
func (s *UserPostgresStorage) bootstrap(ctx context.Context) error {
	tx, err := s.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
		create table if not exists users (
			id serial primary key,
			login text not null unique,
			password text not null
		)
	`)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `
		create index if not exists users_login_index on users (login)
	`)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// NewUserStore creates a new UserPostgresStorage.
//
// It takes a pointer to a pgx.Conn as a parameter and returns a pointer to UserPostgresStorage.
func NewUserStore(c *pgx.Conn) *UserPostgresStorage {
	storage := &UserPostgresStorage{
		conn: c,
	}

	err := storage.bootstrap(context.Background())
	if err != nil {
		panic(err)
	}

	return storage
}
