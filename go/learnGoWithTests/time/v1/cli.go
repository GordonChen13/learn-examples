package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type CLI struct {
	playerStore PlayerStore
	in *bufio.Scanner
	alerter BlindAlerter
}

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI {
		playerStore:store,
		in: bufio.NewScanner(in),
		alerter:alerter,
	}
}

func (cli *CLI) PlayPoker()  {
    blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
    blindTime := 0 * time.Second
    for _, blind := range blinds {
    	cli.alerter.ScheduleAlertAt(blindTime, blind)
    	blindTime = blindTime + 10 * time.Minute
	}

	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
