package infrastructure

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Config holds database configuration
type Config struct {
	Path string // Path to SQLite database file
}

// NewDB creates a new database connection
func NewDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", cfg.Path)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
