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

// New default string representation of structure
func (f *File) String() string {
	return fmt.Sprintf("%s, %v", f.Name, f.ID)
}

func NewFile(filename string, blob []byte) (*File, error) {
	fileID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	return &File{
		fileID,
		filename,
		blob,
	}, nil
}
