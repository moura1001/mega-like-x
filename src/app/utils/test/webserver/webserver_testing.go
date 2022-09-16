package utilstestingwebserver

import (
	"testing"

	"github.com/gorilla/websocket"
)

func MustDialWS(t *testing.T, wsURL string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on '%s': '%v'", wsURL, err)
	}

	return ws
}

func WriteWSMessage(t *testing.T, conn *websocket.Conn, message string) {
	t.Helper()

	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection: '%v'", err)
	}
}
