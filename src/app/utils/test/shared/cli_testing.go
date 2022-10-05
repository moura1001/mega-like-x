package utilstesting

import (
	"io"
	"testing"
)

type PollSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
	FinishCalled bool
}

func (p *PollSpy) Start(numberOfVotingOptions int, to io.Writer) {
	p.StartedWith = numberOfVotingOptions
	p.StartCalled = true
}

func (p *PollSpy) Finish(winner string) {
	p.FinishedWith = winner
	p.FinishCalled = true
}

func AssertGameStartedWith(t *testing.T, poll *PollSpy, want int) {
	t.Helper()

	if poll.StartedWith != want {
		t.Errorf("wanted Start called with %d options but got %d", want, poll.StartedWith)
	}
}

func AssertGameFinishCalledWith(t *testing.T, poll *PollSpy, want string) {
	t.Helper()

	if poll.FinishedWith != want {
		t.Errorf("wanted Finish winner '%s' but got '%s'", want, poll.FinishedWith)
	}
}
