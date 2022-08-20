package cli

import (
	"bufio"
	"fmt"
	"io"
	"moura1001/mega_like_x/src/app/poll"
	apputils "moura1001/mega_like_x/src/app/utils/app"
	"regexp"
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
	winner := extractWinner(userInput)
	if winner == "" {
		fmt.Fprintf(cli.out, apputils.BadWinnerInputErrMsg)
		return
	}

	cli.poll.Finish(winner)
}

func extractWinner(userInput string) string {
	suffix := " wins"
	rgx := fmt.Sprintf(`.+%s$`, suffix)

	matched, _ := regexp.MatchString(rgx, userInput)
	if matched {
		return strings.TrimSuffix(userInput, suffix)
	}

	return ""
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
