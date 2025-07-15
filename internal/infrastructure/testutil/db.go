package testutil

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// NewTestDB creates a new in-memory SQLite database for testing.
// The database will be automatically closed when the test completes.
func NewTestDB(t *testing.T) *sqlx.DB {
	t.Helper()

	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	// Ensure database is closed after test
	t.Cleanup(func() {
		db.Close()
	})

	return db
}

// RunMigrations runs all database migrations.
// This should be called after NewTestDB to set up the schema.
func RunMigrations(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("failed to create sqlite driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		MigrationPath(),
		"sqlite3",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func MigrationPath() string {
	// Find project root using a marker file
	wd, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			break
		}
		parent := filepath.Dir(wd)
		if parent == wd {
			break
		}
		wd = parent
	}
	// The path to migrations must be relative to where the code runs,
	// which differs for `go test ./...` and tests run from VSCode.
	return "file://" + filepath.Join(wd, "migrations")
}
