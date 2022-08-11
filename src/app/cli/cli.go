package cli

import (
	"moura1001/mega_like_x/src/app/store"
	"os"
)

type CLI struct {
	store store.GameStore
}

func NewCLI(storeType store.StoreType, fileDB *os.File) (*CLI, error) {

	cli := new(CLI)

	return cli, nil
}

func (cli *CLI) StartPoll() {
	cli.store.RecordLike("x1")
}
