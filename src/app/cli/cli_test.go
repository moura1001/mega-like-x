package cli

import (
	"moura1001/mega_like_x/src/app/store"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("x1 like\n")
	gameStore := store.StubGameStore{}
	cli, _ := NewCLI("", in, nil)
	cli.store = &gameStore

	cli.StartPoll()

	store.AssertGameLike(t, &gameStore, "x1")
}
