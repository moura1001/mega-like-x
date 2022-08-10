package store

import (
	"encoding/json"
	"fmt"
	"moura1001/mega_like_x/src/app/model"
	"os"
	"sort"
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
	database *json.Encoder
	polling  model.Polling
}

func NewFileSystemGameStore(database *os.File) (*FileSystemGameStore, error) {

	err := initialzeGameDBFile(database)
	if err != nil {
		return nil, fmt.Errorf("problem initializing game db file: %v", err)
	}

	polling, err := model.NewGamePolling(database)

	if err != nil {
		return nil, fmt.Errorf("problem loading game store from file '%s': %v", database.Name(), err)
	}

	return &FileSystemGameStore{
		database: json.NewEncoder(&tape{database}),
		polling:  polling,
	}, nil
}

func initialzeGameDBFile(file *os.File) error {
	file.Seek(0, 0)

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file '%s': %v", file.Name(), err)
	}

	if fileInfo.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

func (f *FileSystemGameStore) GetPolling() model.Polling {
	sort.Slice(f.polling, func(i, j int) bool {
		return f.polling[i].Likes > f.polling[j].Likes
	})
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

	f.database.Encode(f.polling)
}
