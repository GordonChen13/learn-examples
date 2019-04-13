package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
const BadWinnerInputErrMsg = "invalid winner input, expect format of 'PlayerName wins'"

type CLI struct {
	playerStore PlayerStore
	in *bufio.Scanner
	out io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI {
		in: bufio.NewScanner(in),
		out:out,
		game:game,
	}
}

func (cli *CLI) PlayPoker()  {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayers, err := strconv.Atoi(cli.readLine())
	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner, err := extractWinner(winnerInput)

	if err != nil {
		fmt.Fprint(cli.out, BadWinnerInputErrMsg)
		return
	}

	cli.game.Finish(winner)

}

func (c *CLI) scheduleBlindAlerts(numberOfPlayers int) {
	c.game.Start(numberOfPlayers)
}

func extractWinner(userInput string) (string, error){
	if !strings.Contains(userInput, "wins") {
		return "", errors.New(BadWinnerInputErrMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
