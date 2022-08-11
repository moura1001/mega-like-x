package cli

import (
	"moura1001/mega_like_x/src/app/webserver"
	"testing"
)

func TestCLI(t *testing.T) {
	gameStore := webserver.StubGameStore{}
	cli, _ := NewCLI("", nil)
	cli.store = &gameStore

	cli.StartPoll()

	if len(gameStore.GetLikeCalls()) != 1 {
		t.Fatalf("expected a like call but didn't get any")
	}
}
