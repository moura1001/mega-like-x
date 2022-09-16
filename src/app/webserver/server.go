package webserver

import (
	"encoding/json"
	"fmt"
	"html/template"
	"moura1001/mega_like_x/src/app/store"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type GameServer struct {
	store store.GameStore
	http.Handler
	htmlTemplatePath string
}

func NewGameServer(store store.GameStore, htmlTemplatePath string) *GameServer {

	server := new(GameServer)

	server.store = store

	router := mux.NewRouter()
	router.HandleFunc("/games", server.gamesHandler)
	router.HandleFunc("/likes/{game}", server.likesHandler)
	router.HandleFunc("/poll", server.pollHandler)
	router.Handle("/ws", http.HandlerFunc(server.webSocket))

	server.Handler = router

	server.htmlTemplatePath = htmlTemplatePath

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

func (g *GameServer) pollHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(g.htmlTemplatePath)

	if err != nil {
		http.Error(w, fmt.Sprintf("problem loading template: '%s'", err.Error()), http.StatusInternalServerError)
	}

	tmpl.Execute(w, nil)
}

func (g *GameServer) webSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, _ := upgrader.Upgrade(w, r, nil)
	_, winnerMsg, _ := conn.ReadMessage()

	g.store.RecordLike(string(winnerMsg))
}
