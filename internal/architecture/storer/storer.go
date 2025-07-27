package storer

import (
	"fmt"
	"tg-bookmarks-bot/internal/infrastructure/storage"
)

type Storer struct {
	storage storage.Storage
}

func New() (*Storer, error) {
	storer, err := storage.New(storage.StorageTypeSqlite)
	if err != nil {
		return nil, fmt.Errorf("could not create storer: %w", err)
	}

	return &Storer{
		storer,
	}, nil
}
