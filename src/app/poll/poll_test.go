package poll_test

import (
	"moura1001/mega_like_x/src/app/poll"
	utilstesting "moura1001/mega_like_x/src/app/utils/test/shared"
	"testing"
	"time"
)

func TestPollStart(t *testing.T) {

	t.Run("schedules alerts on poll start for 5 players", func(t *testing.T) {
		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		blindAlerter := &utilstesting.SpyBlindAlerter{}
		poll := poll.NewPoll(&gameStore, blindAlerter)

		poll.Start(5)

		cases := []utilstesting.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 400},
			{At: 30 * time.Minute, Amount: 800},
			{At: 40 * time.Minute, Amount: 1600},
		}

		utilstesting.CheckSchedulingCases(t, cases, blindAlerter)
	})

	t.Run("schedules alerts on poll start for 6 players", func(t *testing.T) {
		blindAlerter := &utilstesting.SpyBlindAlerter{}
		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		poll := poll.NewPoll(&gameStore, blindAlerter)

		poll.Start(6)

		cases := []utilstesting.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100},
			{At: 11 * time.Minute, Amount: 200},
			{At: 22 * time.Minute, Amount: 400},
			{At: 33 * time.Minute, Amount: 800},
		}

		utilstesting.CheckSchedulingCases(t, cases, blindAlerter)
	})

}

func TestPollFinish(t *testing.T) {
	winner := "x1"

	gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
	blindAlerter := &utilstesting.SpyBlindAlerter{}
	poll := poll.NewPoll(&gameStore, blindAlerter)

	poll.Finish(winner)
	utilstesting.AssertGameLike(t, &gameStore, winner)
}
