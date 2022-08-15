package main

import (
	"fmt"
	"log"
	"moura1001/mega_like_x/src/app/alerter"
	"moura1001/mega_like_x/src/app/cli"
	"moura1001/mega_like_x/src/app/store"
	"os"
)

func main() {
	fmt.Println("Where's the floor?")
	fmt.Println("Type '{Game} wins' to record a win")

	poll, err := cli.NewCLI(
		store.IN_MEMORY, os.Stdin, nil,
		alerter.BlindAlerterFunc(alerter.StdOutAlerter),
	)

	if err != nil {
		log.Fatalf("poll initialization error: %v", err)
	}

	poll.StartPoll()
}
