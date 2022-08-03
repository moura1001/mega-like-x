package app

import (
	"fmt"
	"moura1001/mega_like_x/src/app/store"
	"net/http"
)

type GameServer struct {
	store store.GameStore
}

func NewGameServer(storeType string) *GameServer {
	server := &GameServer{}

	switch storeType {
	case "in_memory":
		server.store = &store.InMemoryGameStore{}
	}

	return server
}

func (g *GameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	game := r.URL.Path[len("/likes/"):]

	fmt.Fprint(w, g.store.GetGameLikes(game))
}
