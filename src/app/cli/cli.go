package cli

import (
	"bufio"
	"fmt"
	"io"
	"moura1001/mega_like_x/src/app/alerter"
	"moura1001/mega_like_x/src/app/store"
	apputils "moura1001/mega_like_x/src/app/utils/app"
	"os"
	"strconv"
	"strings"
	"time"
)

type CLI struct {
	store   store.GameStore
	in      *bufio.Scanner
	out     io.Writer
	alerter alerter.BlindAlerter
}

func NewCLI(storeType store.StoreType,
	userIn io.Reader, sysOut io.Writer,
	fileDB *os.File, alerter alerter.BlindAlerter,
) (*CLI, error) {

	var err error = nil

	cli := new(CLI)

	cli.in = bufio.NewScanner(userIn)
	cli.out = sysOut
	cli.store, err = store.GetNewGameStore(storeType, fileDB)
	cli.alerter = alerter

	return cli, err
}

func (cli *CLI) StartPoll() {
	fmt.Fprintf(cli.out, apputils.UserPrompt)

	numberOfVotingOptionsInput, _ := strconv.Atoi(cli.readLine())

	cli.scheduleBlindAlerts(numberOfVotingOptionsInput)

	userInput := cli.readLine()
	cli.store.RecordLike(extractVote(userInput))
}

func (cli *CLI) scheduleBlindAlerts(numberOfVotingOptionsInput int) {
	blindIncrement := time.Duration(5+numberOfVotingOptionsInput) * time.Minute

	blinds := []int{100, 200, 400, 800, 1600}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime += blindIncrement
	}
}

func extractVote(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
