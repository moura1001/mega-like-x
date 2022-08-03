package app

import (
	"fmt"
	"net/http"
)

type GameServer struct {
	store GameStore
}

func (g *GameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	game := r.URL.Path[len("/likes/"):]

	fmt.Fprint(w, g.store.GetGameLikes(game))
}

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
