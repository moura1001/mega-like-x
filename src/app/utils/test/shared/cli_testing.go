package utilstesting

type PollSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
	FinishCalled bool
}

func (p *PollSpy) Start(numberOfVotingOptions int) {
	p.StartedWith = numberOfVotingOptions
	p.StartCalled = true
}

func (p *PollSpy) Finish(winner string) {
	p.FinishedWith = winner
	p.FinishCalled = true
}