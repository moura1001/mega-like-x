package cli

import (
	"io"
	"moura1001/mega_like_x/src/app/store"
	"os"
)

type CLI struct {
	store store.GameStore
	in    io.Reader
}

func NewCLI(storeType store.StoreType, userIn io.Reader, fileDB *os.File) (*CLI, error) {

	cli := new(CLI)

	return cli, nil
}

func (cli *CLI) StartPoll() {
	cli.store.RecordLike("x1")
}
