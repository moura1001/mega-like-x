package store

import (
	"io"
	"moura1001/mega_like_x/src/app/model"
)

type FileSystemGameStore struct {
	database io.ReadSeeker
}

func NewFileSystemGameStore(database io.ReadSeeker) *FileSystemGameStore {
	return &FileSystemGameStore{database}
}

func (f *FileSystemGameStore) GetPolling() []model.Game {
	f.database.Seek(0, 0)
	polling, _ := model.NewGamePolling(f.database)
	return polling
}
