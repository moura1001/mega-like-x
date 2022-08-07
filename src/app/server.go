package app

import (
	"encoding/json"
	"fmt"
	"io"
	"moura1001/mega_like_x/src/app/store"
	"net/http"

	"github.com/gorilla/mux"
)

type GameServer struct {
	store store.GameStore
	http.Handler
}

func NewGameServer(storeType store.StoreType, fileDB io.ReadWriteSeeker) *GameServer {
	server := new(GameServer)

	switch storeType {
	case store.IN_MEMORY:
		server.store = store.NewInMemoryGameStore()
	case store.POSTGRES:
		server.store = store.NewPostgresGameStore()
	case store.FILE_SYSTEM:
		server.store = store.NewFileSystemGameStore(fileDB)
	}

	router := mux.NewRouter()
	router.HandleFunc("/games", server.gamesHandler)
	router.HandleFunc("/likes/{game}", server.likesHandler)

	server.Handler = router

	return server
}

func (g *GameServer) gamesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(g.store.GetPolling())
	w.WriteHeader(http.StatusOK)
}

func (g *GameServer) likesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game := vars["game"]

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
