package app

import (
	"fmt"
	"moura1001/mega_like_x/src/app/store"
	"net/http"
)

type GameServer struct {
	store store.GameStore
	http.Handler
}

func NewGameServer(storeType store.StoreType) *GameServer {
	server := new(GameServer)

	switch storeType {
	case store.IN_MEMORY:
		server.store = store.NewInMemoryGameStore()
	case store.POSTGRES:
		server.store = store.NewPostgresGameStore()
	}

	router := http.NewServeMux()
	router.Handle("/games", http.HandlerFunc(server.gamesHandler))
	router.Handle("/likes/", http.HandlerFunc(server.likesHandler))

	server.Handler = router

	return server
}

func (g *GameServer) gamesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (g *GameServer) likesHandler(w http.ResponseWriter, r *http.Request) {
	game := r.URL.Path[len("/likes/"):]

	switch r.Method {
	case http.MethodGet:
		g.showLikes(w, game)
	case http.MethodPost:
		g.processLike(w, game)
	}
}

func (g *GameServer) showLikes(w http.ResponseWriter, game string) {
	likes := g.store.GetGameLikes(game)

	if likes == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, likes)
}

func (g *GameServer) processLike(w http.ResponseWriter, game string) {
	g.store.RecordLike(game)
	w.WriteHeader(http.StatusAccepted)
}
