package cli_test

import (
	"bytes"
	"moura1001/mega_like_x/src/app/cli"
	apputils "moura1001/mega_like_x/src/app/utils/app"
	utilstesting "moura1001/mega_like_x/src/app/utils/test/shared"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("it prompts the user to enter the number of voting options and starts the poll", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("6\n")
		poll := &utilstesting.PollSpy{}

		c := cli.NewCLI(in, stdout, poll)

		c.StartPoll()

		gotPrompt := stdout.String()
		wantPrompt := apputils.UserPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if poll.StartedWith != 6 {
			t.Errorf("wanted Start called with 6 options but got %d", poll.StartedWith)
		}
	})

}
