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

	t.Run("start poll with 4 game options and finish with 'x2' as winner", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		poll := &utilstesting.PollSpy{}

		in := userSends("4", "x2 wins")

		c := cli.NewCLI(in, stdout, poll)

		c.StartPoll()

		assertMessageSentToUser(t, stdout, apputils.UserPrompt)
		assertGameStartedWith(t, poll, 4)
		assertGameFinishCalledWith(t, poll, "x2")
	})

	t.Run("start poll with 6 game options and finish with 'x6' as winner", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		poll := &utilstesting.PollSpy{}

		in := userSends("6", "x6 wins")

		c := cli.NewCLI(in, stdout, poll)

		c.StartPoll()

		assertMessageSentToUser(t, stdout, apputils.UserPrompt)
		assertGameStartedWith(t, poll, 6)
		assertGameFinishCalledWith(t, poll, "x6")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the poll", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		poll := &utilstesting.PollSpy{}

		in := userSends("x7")

		c := cli.NewCLI(in, stdout, poll)

		c.StartPoll()

		assertPollNotStarted(t, poll)
		assertMessageSentToUser(t, stdout, apputils.UserPrompt, apputils.BadUserInputErrMsg)

	})

	t.Run("start poll with 8 game options and prints an error when invalid winner input is entered and does not finish the poll", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		poll := &utilstesting.PollSpy{}

		in := userSends("8", "x6 is better than x4")

		c := cli.NewCLI(in, stdout, poll)

		c.StartPoll()

		if poll.FinishCalled {
			t.Errorf("poll should not have finished")
		}

		assertMessageSentToUser(t, stdout, apputils.UserPrompt, apputils.BadWinnerInputErrMsg)
	})

	t.Run("start poll with 3 game options and finish with hypothetical game title 'Wins to wins' as winner", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		poll := &utilstesting.PollSpy{}

		in := userSends("3", "Wins to wins wins")

		c := cli.NewCLI(in, stdout, poll)

		c.StartPoll()

		assertMessageSentToUser(t, stdout, apputils.UserPrompt)
		assertGameStartedWith(t, poll, 3)
		assertGameFinishCalledWith(t, poll, "Wins to wins")
	})

}

func userSends(inputs ...string) *strings.Reader {

	in := strings.Join(inputs, "\n")

	return strings.NewReader(in)
}

func assertMessageSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()

	got := stdout.String()
	want := strings.Join(messages, "")

	if got != want {
		t.Errorf("got '%s' sent to stdout but expected %+v", got, messages)
	}
}

func assertGameStartedWith(t *testing.T, poll *utilstesting.PollSpy, want int) {
	t.Helper()

	if poll.StartedWith != want {
		t.Errorf("wanted Start called with %d options but got %d", want, poll.StartedWith)
	}
}

func assertGameFinishCalledWith(t *testing.T, poll *utilstesting.PollSpy, want string) {
	t.Helper()

	if poll.FinishedWith != want {
		t.Errorf("wanted Finish winner '%s' but got '%s'", want, poll.FinishedWith)
	}
}

func assertPollNotStarted(t *testing.T, poll *utilstesting.PollSpy) {
	t.Helper()

	if poll.StartCalled {
		t.Errorf("poll should not have started")
	}
}
