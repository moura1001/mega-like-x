package store

import (
	"io"
	"moura1001/mega_like_x/src/app/model"
)

type FileSystemGameStore struct {
	database io.Reader
}

func NewFileSystemGameStore(database io.Reader) *FileSystemGameStore {
	return &FileSystemGameStore{database}
}

func (f *FileSystemGameStore) GetPolling() []model.Game {
	polling, _ := model.NewGamePolling(f.database)
	return polling
}
