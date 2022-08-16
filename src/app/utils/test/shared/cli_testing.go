package utilstesting

type PollSpy struct {
	StartedWith  int
	FinishedWith string
}

func (p *PollSpy) Start(numberOfVotingOptions int) {
	p.StartedWith = numberOfVotingOptions
}

func (p *PollSpy) Finish(winner string) {
	p.FinishedWith = winner
}
