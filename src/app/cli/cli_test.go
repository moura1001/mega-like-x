package cli

import (
	"fmt"
	utilstesting "moura1001/mega_like_x/src/app/utils/test/shared"
	"strings"
	"testing"
	"time"
)

type SpyBlindAlerter struct {
	alerts []struct {
		scheduledAt time.Duration
		amount      int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, struct {
		scheduledAt time.Duration
		amount      int
	}{duration, amount})
}

var dummySpyAlerter = &SpyBlindAlerter{}

func TestCLI(t *testing.T) {

	t.Run("record x1 win from user input", func(t *testing.T) {
		in := strings.NewReader("x1 wins\n")
		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		cli, _ := NewCLI("", in, nil, dummySpyAlerter)
		cli.store = &gameStore

		cli.StartPoll()

		utilstesting.AssertGameLike(t, &gameStore, "x1")
	})

	t.Run("record x6 win from user input", func(t *testing.T) {
		in := strings.NewReader("x6 wins\n")
		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		cli, _ := NewCLI("", in, nil, dummySpyAlerter)
		cli.store = &gameStore

		cli.StartPoll()

		utilstesting.AssertGameLike(t, &gameStore, "x6")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("x2 wins\n")
		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		blindAlerter := &SpyBlindAlerter{}

		cli, _ := NewCLI("", in, nil, blindAlerter)
		cli.store = &gameStore

		cli.StartPoll()

		cases := []struct {
			expectedScheduleTime time.Duration
			expectedAmount       int
		}{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 400},
			{30 * time.Minute, 800},
			{40 * time.Minute, 1600},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", c.expectedAmount, c.expectedScheduleTime), func(t *testing.T) {

				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled, got: %v", i, blindAlerter.alerts)
				}

				alert := blindAlerter.alerts[i]

				gotAmount := alert.amount
				if gotAmount != c.expectedAmount {
					t.Errorf("got amount %d, want %d", gotAmount, c.expectedAmount)
				}

				gotScheduledTime := alert.scheduledAt
				if gotScheduledTime != c.expectedScheduleTime {
					t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, c.expectedScheduleTime)
				}
			})
		}
	})

}
