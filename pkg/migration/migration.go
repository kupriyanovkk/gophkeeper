package migration

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kupriyanovkk/gophkeeper/internal/server/config"
)

type Migration struct {
	config config.Config
}

// Up performs the migration Up operation.
//
// It returns an error.
func (m *Migration) Up() error {
	migration, err := migrate.New("file://migrations", m.config.DatabaseDSN)
	if err != nil {
		return fmt.Errorf("migrate.New: %w", err)
	}

	defer migration.Close()

	return migration.Up()
}

// DropTest drops the test database.
//
// No parameters.
// Error.
func (m *Migration) DropTest() error {
	migration, err := migrate.New("file://../../migrations", m.config.DatabaseDSNTest)
	if err != nil {
		return fmt.Errorf("mingrate.New error: %w", err)
	}
	defer migration.Close()

	if err = migration.Drop(); err != nil {
		return fmt.Errorf("error in migrating db: %w", err)
	}

	return nil
}

// UpTest is a function that performs the database migration for testing environment.
//
// It does not take any parameters and returns an error.
func (m *Migration) UpTest() error {
	migration, err := migrate.New("file://../../internal/server/migrations", m.config.DatabaseDSNTest)
	if err != nil {
		return fmt.Errorf("mingrate.New error: %w", err)
	}
	defer migration.Close()

	if err = migration.Up(); err != nil {
		return fmt.Errorf("error in migrating db: %w", err)
	}

	return nil
}

// RefreshTest is a Go function that drops and then updates a test, returning an error if there is one.
//
// None. Error.
func (m *Migration) RefreshTest() error {
	if err := m.DropTest(); err != nil {
		return err
	}
	if err := m.UpTest(); err != nil {
		return err
	}
	return nil
}

// NewMigration creates a new Migration with the given config.
//
// It takes a config.Config parameter and returns a *Migration pointer.
func NewMigration(config config.Config) *Migration {
	return &Migration{config: config}
}
