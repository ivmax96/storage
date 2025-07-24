package file

import (
	"fmt"

	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) *File {
	return &File{
		ID:   uuid.New(),
		Name: filename,
		Data: blob,
	}
}

func (f *File) String() string {
	return fmt.Sprintf("file(%s, %v)", f.Name, f.ID)
}
