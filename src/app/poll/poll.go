package poll

import (
	"moura1001/mega_like_x/src/app/alerter"
	"moura1001/mega_like_x/src/app/store"
	"time"
)

type Poll struct {
	store   store.GameStore
	alerter alerter.BlindAlerter
}

func NewPoll(store store.GameStore, alerter alerter.BlindAlerter) *Poll {

	return &Poll{
		store:   store,
		alerter: alerter,
	}
}

func (p *Poll) Start(numberOfVotingOptions int) {
	blindIncrement := time.Duration(5+numberOfVotingOptions) * time.Minute

	blinds := []int{100, 200, 400, 800, 1600}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime += blindIncrement
	}
}

func (p *Poll) Finish(winner string) {
	p.store.RecordLike(winner)
}
