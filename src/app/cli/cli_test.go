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

	t.Run("it prints an error when a non numeric value is entered and does not start the poll", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("x7\n")
		poll := &utilstesting.PollSpy{}

		c := cli.NewCLI(in, stdout, poll)

		c.StartPoll()

		if poll.StartCalled {
			t.Errorf("poll should not have started")
		}

		gotPrompt := stdout.String()
		wantPrompt := apputils.UserPrompt + apputils.BadUserInputErrMsg

		if gotPrompt != wantPrompt {
			t.Errorf("got '%s', want '%s'", gotPrompt, wantPrompt)
		}
	})

}
