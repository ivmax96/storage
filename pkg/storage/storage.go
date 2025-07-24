package storage

import "github.com/ivmax96/storage/internal/storage"

func New() *storage.Storage {
	return storage.NewStorage()
}
