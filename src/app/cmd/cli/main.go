package main

import (
	"fmt"
	"moura1001/mega_like_x/src/app/alerter"
	"moura1001/mega_like_x/src/app/cli"
	"moura1001/mega_like_x/src/app/poll"
	"moura1001/mega_like_x/src/app/store"
	"os"
)

func main() {
	fmt.Println("Where's the floor?")
	fmt.Println("Type '{Game} wins' to record a win")

	store := store.NewInMemoryGameStore()
	poll := poll.NewMegaLike(
		store,
		alerter.BlindAlerterFunc(alerter.StdOutAlerter),
	)

	cli := cli.NewCLI(os.Stdin, nil, poll)

	cli.StartPoll()
}
