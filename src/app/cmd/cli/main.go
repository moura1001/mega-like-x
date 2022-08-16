package main

import (
	"fmt"
	"moura1001/mega_like_x/src/app/alerter"
	"moura1001/mega_like_x/src/app/cli"
	"moura1001/mega_like_x/src/app/store"
	"os"
)

func main() {
	fmt.Println("Where's the floor?")
	fmt.Println("Type '{Game} wins' to record a win")

	store := store.NewInMemoryGameStore()
	poll := cli.NewCLI(
		store, os.Stdin, nil,
		alerter.BlindAlerterFunc(alerter.StdOutAlerter),
	)

	poll.StartPoll()
}
