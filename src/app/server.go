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

	switch r.Method {
	case http.MethodGet:
		g.showLikes(w, r)
	case http.MethodPost:
		g.processLike(w)
	}
}

func (g *GameServer) showLikes(w http.ResponseWriter, r *http.Request) {
	game := r.URL.Path[len("/likes/"):]

	likes := g.store.GetGameLikes(game)

	if likes == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, likes)
}

func (g *GameServer) processLike(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}
