package poll

import (
	"moura1001/mega_like_x/src/app/alerter"
	"moura1001/mega_like_x/src/app/store"
	"time"
)

type Poll interface {
	Start(numberOfVotingOptions int)
	Finish(winner string)
}

type MegaLike struct {
	store   store.GameStore
	alerter alerter.BlindAlerter
}

func NewMegaLike(store store.GameStore, alerter alerter.BlindAlerter) *MegaLike {

	return &MegaLike{
		store:   store,
		alerter: alerter,
	}
}

func (p *MegaLike) Start(numberOfVotingOptions int) {
	blindIncrement := time.Duration(5+numberOfVotingOptions) * time.Minute

	blinds := []int{100, 200, 400, 800, 1600}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime += blindIncrement
	}
}

func (p *MegaLike) Finish(winner string) {
	p.store.RecordLike(winner)
}
