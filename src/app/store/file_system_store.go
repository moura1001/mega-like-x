package store

import (
	"encoding/json"
	"io"
	"moura1001/mega_like_x/src/app/model"
	"os"
)

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)
	return t.file.Write(p)
}

type FileSystemGameStore struct {
	database io.Writer
	polling  model.Polling
}

func NewFileSystemGameStore(database *os.File) *FileSystemGameStore {
	database.Seek(0, 0)
	polling, _ := model.NewGamePolling(database)

	return &FileSystemGameStore{
		database: &tape{database},
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

	json.NewEncoder(f.database).Encode(f.polling)
}
