package store

import (
	"moura1001/mega_like_x/src/app/model"
)

type GameStore interface {
	GetGameLikes(name string) int
	RecordLike(name string)
	GetPolling() model.Polling
}
