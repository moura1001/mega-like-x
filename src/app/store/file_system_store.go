package store

import (
	"encoding/json"
	"io"
	"moura1001/mega_like_x/src/app/model"
)

type FileSystemGameStore struct {
	database io.ReadWriteSeeker
}

func NewFileSystemGameStore(database io.ReadWriteSeeker) *FileSystemGameStore {
	return &FileSystemGameStore{database}
}

func (f *FileSystemGameStore) GetPolling() model.Polling {
	f.database.Seek(0, 0)
	polling, _ := model.NewGamePolling(f.database)
	return polling
}

func (f *FileSystemGameStore) GetGameLikes(name string) int {
	game := f.GetPolling().Find(name)

	if game != nil {
		return game.Likes
	}

	return 0
}

func (f *FileSystemGameStore) RecordLike(name string) {
	polling := f.GetPolling()
	game := polling.Find(name)

	if game != nil {
		game.Likes++
	}

	f.database.Seek(0, 0)

	json.NewEncoder(f.database).Encode(polling)
}
