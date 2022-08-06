package store

import "moura1001/mega_like_x/src/app/model"

type GameStore interface {
	GetGameLikes(name string) int
	RecordLike(name string)
	GetPolling() []model.Game
}

type StoreType string

const (
	IN_MEMORY StoreType = "in_memory"
	POSTGRES  StoreType = "postgres"
)
