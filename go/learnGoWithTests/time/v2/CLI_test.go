package poker

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

type scheduledAlert struct {
	at time.Duration
	amount int
}
type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}

var dummySpyAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("7\nChris wins\n")
		playerStore := &StubPlayerStore{}

		game := NewGame(dummySpyAlerter, playerStore)
		cli := NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		want := "Chris"

		assertPlayerWin(t, playerStore, want)
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("7\nCleo wins\n")
		playerStore := &StubPlayerStore{}

		game := NewGame(dummySpyAlerter, playerStore)
		cli := NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		want := "Cleo"

		assertPlayerWin(t, playerStore, want)
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		blindAlerter := &SpyBlindAlerter{}

		game := NewGame(blindAlerter, dummyPlayerStore)
		cli := NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		if len(blindAlerter.alerts) < 1 {
			t.Fatal("expected a blind alert to be scheduled")
		}
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		blindAlerter := &SpyBlindAlerter{}

		game := NewGame(blindAlerter, dummyPlayerStore)
		cli := NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		cases := []scheduledAlert{
			{0 * time.Minute, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want:= range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {

				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
        stdout := &bytes.Buffer{}

		game := NewGame(dummySpyAlerter, dummyPlayerStore)
		cli := NewCLI(dummyStdIn, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := PlayerPrompt
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		blindAlerter := &SpyBlindAlerter{}

		game := NewGame(blindAlerter, dummyPlayerStore)
		cli := NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := PlayerPrompt

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

		cases := []scheduledAlert{
			{0 * time.Minute, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		for i, want:= range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {

				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})

}

func assertScheduledAlert(t *testing.T, got,want scheduledAlert)  {
	if got.amount!= want.amount{
		t.Errorf("got amount %d, want %d",got.amount, want.amount)
	}

	if got.at != want.at {
		t.Errorf("got scheduled time of %v, want %v", got.at, want.at)
	}

}

func assertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	if store.winCalls[0] != winner {
		t.Errorf("didn't record correct winner, got '%s', want '%s'", store.winCalls[0], winner)
	}

}
