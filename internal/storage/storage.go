package storage

import (
	"fmt"
	"github.com/Kiatsyndesi/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	Files map[uuid.UUID]*file.File
}

// Constructor in go style
func NewStorage() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, data []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, data)

	if err != nil {
		return nil, err
	}

	s.Files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.Files[fileID]

	if !ok {
		return nil, fmt.Errorf("file with %v ID was not found", fileID)
	}

	return foundFile, nil
}
