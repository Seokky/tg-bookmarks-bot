package storage

import (
	"fmt"
	"log"
	"tg-bookmarks-bot/internal/infrastructure/storage/internal/sqlite"
)

type Storage interface {
	Insert(username, bookmark string) error
	Random(username string) (string, error)
	Count(username string) (int, error)
	Delete(username, bookmark string) error
}

const (
	StorageTypeSqlite = "sqlite"
)

func New(storageType string) (Storage, error) {
	switch storageType {
	case StorageTypeSqlite:
		storage, err := sqlite.New()
		if err != nil {
			log.Fatal(err)
		}
		return storage, nil
	}
	return nil, fmt.Errorf("unknown storage type: %s", storageType)
}
