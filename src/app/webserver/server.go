package webserver

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"moura1001/mega_like_x/src/app/poll"
	"moura1001/mega_like_x/src/app/store"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type GameServer struct {
	store store.GameStore
	http.Handler
	htmlTemplate *template.Template
	poll         poll.Poll
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewGameServer(store store.GameStore, htmlTemplatePath string, poll poll.Poll) (*GameServer, error) {

	server := new(GameServer)

	tmpl, err := template.ParseFiles(htmlTemplatePath)
	if htmlTemplatePath != "" && err != nil {
		return nil, fmt.Errorf("problem opening template '%s': '%v'", htmlTemplatePath, err)
	}

	server.htmlTemplate = tmpl
	server.store = store
	server.poll = poll

	router := mux.NewRouter()
	router.HandleFunc("/games", server.gamesHandler)
	router.HandleFunc("/likes/{game}", server.likesHandler)
	router.HandleFunc("/poll", server.pollHandler)
	router.Handle("/ws", http.HandlerFunc(server.webSocket))

	server.Handler = router

	return server, nil
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

	g.htmlTemplate.Execute(w, nil)
}

func (g *GameServer) webSocket(w http.ResponseWriter, r *http.Request) {

	conn, _ := wsUpgrader.Upgrade(w, r, nil)

	_, numberOfVotingOptionsMsg, _ := conn.ReadMessage()
	numberOfVotingOptions, _ := strconv.Atoi(string(numberOfVotingOptionsMsg))
	// TODO: don't discard the blinds messages
	g.poll.Start(numberOfVotingOptions, io.Discard)

	_, winnerMsg, _ := conn.ReadMessage()

	g.poll.Finish(string(winnerMsg))
}
