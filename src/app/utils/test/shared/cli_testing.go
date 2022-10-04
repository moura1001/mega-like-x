package utilstesting

import "io"

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
