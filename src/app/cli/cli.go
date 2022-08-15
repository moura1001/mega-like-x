package cli

import (
	"bufio"
	"io"
	"moura1001/mega_like_x/src/app/alerter"
	"moura1001/mega_like_x/src/app/store"
	"os"
	"strings"
	"time"
)

type CLI struct {
	store   store.GameStore
	in      *bufio.Scanner
	alerter alerter.BlindAlerter
}

func NewCLI(storeType store.StoreType, userIn io.Reader,
	fileDB *os.File, alerter alerter.BlindAlerter,
) (*CLI, error) {

	var err error = nil

	cli := new(CLI)

	cli.in = bufio.NewScanner(userIn)
	cli.store, err = store.GetNewGameStore(storeType, fileDB)
	cli.alerter = alerter

	return cli, err
}

func (cli *CLI) StartPoll() {
	cli.scheduleBlindAlerts()
	userInput := cli.readLine()
	cli.store.RecordLike(extractVote(userInput))
}

func (cli *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 400, 800, 1600}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime += 10 * time.Minute
	}
}

func extractVote(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
