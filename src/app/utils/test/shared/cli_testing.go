package utilstesting

import (
	"io"
	"testing"
	"time"
)

type PollSpy struct {
	StartCalled bool
	StartedWith int
	BlindAlert  []byte

	FinishCalled bool
	FinishedWith string
}

func (p *PollSpy) Start(numberOfVotingOptions int, to io.Writer) {
	p.StartedWith = numberOfVotingOptions
	p.StartCalled = true
	to.Write(p.BlindAlert)
}

func (p *PollSpy) Finish(winner string) {
	p.FinishedWith = winner
	p.FinishCalled = true
}

func AssertGameStartedWith(t *testing.T, poll *PollSpy, want int) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return poll.StartedWith == want
	})

	if !passed {
		t.Errorf("wanted Start called with %d options but got %d", want, poll.StartedWith)
	}
}

func AssertGameFinishCalledWith(t *testing.T, poll *PollSpy, want string) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return poll.FinishedWith == want
	})

	if !passed {
		t.Errorf("wanted Finish winner '%s' but got '%s'", want, poll.FinishedWith)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}
