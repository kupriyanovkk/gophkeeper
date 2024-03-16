package utils

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/kupriyanovkk/gophkeeper/internal/server/config"
	"github.com/kupriyanovkk/gophkeeper/pkg/migration"
)

// CreatePostgresConn creates a PostgreSQL connection.
//
// Takes a context and a connection string as parameters.
// Returns a pgx.Conn pointer.
func CreatePostgresConn(ctx context.Context, connString string) *pgx.Conn {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		panic(err)
	}

	errPing := conn.Ping(ctx)
	if errPing != nil {
		panic(errPing)
	}

	return conn
}

// RefreshTestDatabase refreshes the test database.
//
// It does not take any parameters.
// It does not return anything.
func RefreshTestDatabase() {
	migrator := migration.NewMigration(config.NewConfig())
	if err := migrator.RefreshTest(); err != nil {
		panic(err)
	}
}
