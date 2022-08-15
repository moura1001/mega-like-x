package cli

import (
	utilstesting "moura1001/mega_like_x/src/app/utils/test/shared"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("record x1 win from user input", func(t *testing.T) {
		in := strings.NewReader("x1 wins\n")
		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		cli, _ := NewCLI("", in, nil)
		cli.store = &gameStore

		cli.StartPoll()

		utilstesting.AssertGameLike(t, &gameStore, "x1")
	})

	t.Run("record x6 win from user input", func(t *testing.T) {
		in := strings.NewReader("x6 wins\n")
		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		cli, _ := NewCLI("", in, nil)
		cli.store = &gameStore

		cli.StartPoll()

		utilstesting.AssertGameLike(t, &gameStore, "x6")
	})

}
