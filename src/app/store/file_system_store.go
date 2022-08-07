package store

import (
	"encoding/json"
	"io"
	"moura1001/mega_like_x/src/app/model"
)

type FileSystemGameStore struct {
	database io.ReadWriteSeeker
	polling  model.Polling
}

func NewFileSystemGameStore(database io.ReadWriteSeeker) *FileSystemGameStore {
	database.Seek(0, 0)
	polling, _ := model.NewGamePolling(database)
	return &FileSystemGameStore{
		database: database,
		polling:  polling,
	}
}

func (f *FileSystemGameStore) GetPolling() model.Polling {
	return f.polling
}

func (f *FileSystemGameStore) GetGameLikes(name string) int {
	game := f.polling.Find(name)

	if game != nil {
		return game.Likes
	}

	return 0
}

func (f *FileSystemGameStore) RecordLike(name string) {
	game := f.polling.Find(name)

	if game != nil {
		game.Likes++
	} else {
		f.polling = append(f.polling, model.Game{Name: name, Likes: 1})
	}

	f.database.Seek(0, 0)

	json.NewEncoder(f.database).Encode(f.polling)
}
