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
	in    io.Reader
}

func NewCLI(storeType store.StoreType, userIn io.Reader, fileDB *os.File) (*CLI, error) {

	var err error = nil

	cli := new(CLI)

	cli.in = userIn

	cli.store, err = store.GetNewGameStore(storeType, fileDB)

	return cli, err
}

func (cli *CLI) StartPoll() {
	reader := bufio.NewScanner(cli.in)
	reader.Scan()
	cli.store.RecordLike(extractVote(reader.Text()))
}

func extractVote(userInput string) string {
	return strings.Replace(userInput, " like", "", 1)
}
