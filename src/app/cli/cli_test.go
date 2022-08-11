package cli

import (
	utilstesting "moura1001/mega_like_x/src/app/utils/test/shared"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("record x1 like from user input", func(t *testing.T) {
		in := strings.NewReader("x1 like\n")
		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		cli, _ := NewCLI("", in, nil)
		cli.store = &gameStore

		cli.StartPoll()

		utilstesting.AssertGameLike(t, &gameStore, "x1")
	})

	t.Run("record x6 like from user input", func(t *testing.T) {
		in := strings.NewReader("x6 like\n")
		gameStore := utilstesting.StubGameStore{}
		cli, _ := NewCLI("", in, nil)
		cli.store = &gameStore

		cli.StartPoll()

		utilstesting.AssertGameLike(t, &gameStore, "x6")
	})

}
