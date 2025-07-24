package storage

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ivmax96/storage/internal/file"
)

type Storage struct {
	files map[uuid.UUID]file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]file.File),
	}
}

func (st *Storage) Upload(filename string, blob []byte) *file.File {
	newFile := file.NewFile(filename, blob)

	st.files[newFile.ID] = *newFile

	return newFile
}

func (st *Storage) GetByID(id uuid.UUID) (*file.File, error) {
	f, ok := st.files[id]
	if !ok {
		return nil, fmt.Errorf("file with id %v not found", id)
	}
	return &f, nil
}
