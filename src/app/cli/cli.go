package cli

import (
	"bufio"
	"fmt"
	"io"
	"moura1001/mega_like_x/src/app/poll"
	apputils "moura1001/mega_like_x/src/app/utils/app"
	"strconv"
	"strings"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	poll poll.Poll
}

func NewCLI(userIn io.Reader, sysOut io.Writer, poll poll.Poll) *CLI {

	return &CLI{
		in:   bufio.NewScanner(userIn),
		out:  sysOut,
		poll: poll,
	}
}

func (cli *CLI) StartPoll() {
	fmt.Fprintf(cli.out, apputils.UserPrompt)

	numberOfVotingOptionsInput, err := strconv.Atoi(cli.readLine())
	if err != nil {
		fmt.Fprintf(cli.out, apputils.BadUserInputErrMsg)
		return
	}

	cli.poll.Start(numberOfVotingOptionsInput)

	userInput := cli.readLine()
	cli.poll.Finish(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
