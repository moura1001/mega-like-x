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

func GetGameLikes(name string) string {
	if name == "x1" {
		return "32"
	}

	if name == "x2" {
		return "64"
	}

	return ""
}
