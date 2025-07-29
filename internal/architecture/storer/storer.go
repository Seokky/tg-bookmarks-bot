// Package storer implements abstraction layer to work with storage
package storer

import (
	"fmt"
	"tg-bookmarks-bot/internal/infrastructure/storage"
)

// Storer is abstraction layer to work with low level storage
type Storer struct {
	storage storage.Storage
}

// New creates instance of Storer
func New() (*Storer, error) {
	storer, err := storage.New(storage.KindSqlite)
	if err != nil {
		return nil, fmt.Errorf("could not create storer: %w", err)
	}

	return &Storer{
		storer,
	}, nil
}
