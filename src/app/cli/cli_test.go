package cli_test

import (
	"bytes"
	"moura1001/mega_like_x/src/app/cli"
	"moura1001/mega_like_x/src/app/poll"
	apputils "moura1001/mega_like_x/src/app/utils/app"
	utilstesting "moura1001/mega_like_x/src/app/utils/test/shared"
	"strings"
	"testing"
)

var dummySpyAlerter = &utilstesting.SpyBlindAlerter{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {

	t.Run("record x1 win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nx1 wins\n")
		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		poll := poll.NewMegaLike(&gameStore, dummySpyAlerter)

		cli := cli.NewCLI(in, dummyStdOut, poll)

		cli.StartPoll()

		utilstesting.AssertGameLike(t, &gameStore, "x1")
	})

	t.Run("record x6 win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nx6 wins\n")
		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		poll := poll.NewMegaLike(&gameStore, dummySpyAlerter)

		cli := cli.NewCLI(in, dummyStdOut, poll)

		cli.StartPoll()

		utilstesting.AssertGameLike(t, &gameStore, "x6")
	})

	t.Run("it prompts the user to enter the number of voting options", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("6\n")
		blindAlerter := &utilstesting.SpyBlindAlerter{}
		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		poll := poll.NewMegaLike(&gameStore, blindAlerter)

		c := cli.NewCLI(in, stdout, poll)

		c.StartPoll()

		got := stdout.String()
		want := apputils.UserPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
