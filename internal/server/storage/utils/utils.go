package utils

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/kupriyanovkk/gophkeeper/internal/server/config"
	"github.com/kupriyanovkk/gophkeeper/pkg/migration"
)

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

func RefreshTestDatabase() {
	migrator := migration.NewMigration(config.NewConfig())
	if err := migrator.RefreshTest(); err != nil {
		panic(err)
	}
}
