package store

type GameStore interface {
	GetGameLikes(name string) int
}

func GetGameLikes(name string) string {
	if name == "x1" {
		return "32"
	}

	if name == "x2" {
		return "64"
	}

	return ""
}
