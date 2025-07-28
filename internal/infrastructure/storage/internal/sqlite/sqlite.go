// Package sqlite implements methods to work with SQLite storage
package sqlite

import (
	"database/sql"
	"fmt"
	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

const (
	pathToDb  = "./internal/infrastructure/storage/internal/sqlite/storage.db"
	tableName = "bookmarks"
)

// New creates instance of Storage
func New() (*Storage, error) {
	db, err := sql.Open("sqlite3", pathToDb)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	initialQuery := `CREATE TABLE IF NOT EXISTS ` + tableName + ` (username TEXT, bookmark TEXT)`
	if _, err := db.Exec(initialQuery); err != nil {
		return nil, fmt.Errorf("failed to create initial table: %w", err)
	}

	return &Storage{db}, nil
}
