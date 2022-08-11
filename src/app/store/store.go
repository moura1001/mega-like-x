package store

import (
	"moura1001/mega_like_x/src/app/model"
	"os"
)

type GameStore interface {
	GetGameLikes(name string) int
	RecordLike(name string)
	GetPolling() model.Polling
}

type StoreType string

const (
	IN_MEMORY   StoreType = "in_memory"
	POSTGRES    StoreType = "postgres"
	FILE_SYSTEM StoreType = "file_system"
)

func GetNewGameStore(storeType StoreType, fileDB *os.File) (GameStore, error) {
	var store GameStore = nil
	var err error = nil

	switch storeType {
	case IN_MEMORY:
		store = NewInMemoryGameStore()
	case POSTGRES:
		store = NewPostgresGameStore()
	case FILE_SYSTEM:
		store, err = NewFileSystemGameStore(fileDB)
	}

	return store, err
}
