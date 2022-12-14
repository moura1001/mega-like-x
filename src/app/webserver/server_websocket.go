package webserver

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type pollServerWS struct {
	*websocket.Conn
}

func newPollServerWS(w http.ResponseWriter, r *http.Request) *pollServerWS {
	conn, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("problem upgrading connection to WebSocket: '%v'\n", err)
	}

	return &pollServerWS{
		conn,
	}
}

func (ws *pollServerWS) WaitForMessage() string {

	_, msg, err := ws.ReadMessage()

	if err != nil {
		log.Printf("error reading from WebSocket: '%v'\n", err)
	}

	return string(msg)
}

func (ws *pollServerWS) Write(b []byte) (n int, err error) {
	err = ws.WriteMessage(1, b)

	if err != nil {
		return 0, err
	}

	return len(b), nil
}
