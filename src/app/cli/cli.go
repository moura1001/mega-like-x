package cli

import (
	"bufio"
	"io"
	"moura1001/mega_like_x/src/app/store"
	"os"
	"strings"
)

type CLI struct {
	store store.GameStore
	in    *bufio.Scanner
}

func NewCLI(storeType store.StoreType, userIn io.Reader, fileDB *os.File) (*CLI, error) {

	var err error = nil

	cli := new(CLI)

	cli.in = bufio.NewScanner(userIn)

	cli.store, err = store.GetNewGameStore(storeType, fileDB)

	return cli, err
}

func (cli *CLI) StartPoll() {
	userInput := cli.readLine()
	cli.store.RecordLike(extractVote(userInput))
}

func extractVote(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
