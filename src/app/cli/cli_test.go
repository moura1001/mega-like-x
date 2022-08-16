package cli_test

import (
	"bytes"
	"fmt"
	"moura1001/mega_like_x/src/app/cli"
	apputils "moura1001/mega_like_x/src/app/utils/app"
	utilstesting "moura1001/mega_like_x/src/app/utils/test/shared"
	"strings"
	"testing"
	"time"
)

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d likes multiplier at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}

var dummySpyAlerter = &SpyBlindAlerter{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {

	t.Run("record x1 win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nx1 wins\n")
		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		cli := cli.NewCLI(&gameStore, in, dummyStdOut, dummySpyAlerter)

		cli.StartPoll()

		utilstesting.AssertGameLike(t, &gameStore, "x1")
	})

	t.Run("record x6 win from user input", func(t *testing.T) {
		in := strings.NewReader("5\nx6 wins\n")
		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		cli := cli.NewCLI(&gameStore, in, dummyStdOut, dummySpyAlerter)

		cli.StartPoll()

		utilstesting.AssertGameLike(t, &gameStore, "x6")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("5\nx2 wins\n")
		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		blindAlerter := &SpyBlindAlerter{}

		cli := cli.NewCLI(&gameStore, in, dummyStdOut, blindAlerter)

		cli.StartPoll()

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 400},
			{30 * time.Minute, 800},
			{40 * time.Minute, 1600},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {

				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled, got: %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})

	t.Run("it prompts the user to enter the number of voting options", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("6\n")
		blindAlerter := &SpyBlindAlerter{}

		gameStore := utilstesting.GetNewStubGameStore(nil, nil, nil)
		c := cli.NewCLI(&gameStore, in, stdout, blindAlerter)

		c.StartPoll()

		got := stdout.String()
		want := apputils.UserPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{11 * time.Minute, 200},
			{22 * time.Minute, 400},
			{33 * time.Minute, 800},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {

				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled, got: %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})

}

func assertScheduledAlert(t *testing.T, got, want scheduledAlert) {

	if got.amount != want.amount {
		t.Errorf("got amount %d, want %d", got.amount, want.amount)
	}

	if got.at != want.at {
		t.Errorf("got scheduled time of %v, want %v", got.at, want.at)
	}
}
