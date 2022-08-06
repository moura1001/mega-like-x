package store

type GameStore interface {
	GetGameLikes(name string) int
	RecordLike(name string)
}

type StoreType string

const (
	IN_MEMORY StoreType = "in_memory"
	POSTGRES  StoreType = "postgres"
)
