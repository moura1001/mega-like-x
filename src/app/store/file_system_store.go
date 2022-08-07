package store

import (
	"io"
	"moura1001/mega_like_x/src/app/model"
)

type FileSystemGameStore struct {
	database io.ReadWriteSeeker
}

func NewFileSystemGameStore(database io.ReadWriteSeeker) *FileSystemGameStore {
	return &FileSystemGameStore{database}
}

func (f *FileSystemGameStore) GetPolling() []model.Game {
	f.database.Seek(0, 0)
	polling, _ := model.NewGamePolling(f.database)
	return polling
}

func (f *FileSystemGameStore) GetGameLikes(name string) int {
	var likes int

	for _, game := range f.GetPolling() {
		if game.Name == name {
			likes = game.Likes
			break
		}
	}

	return likes
}
