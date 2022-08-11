package main

import (
	"fmt"
	"log"
	"moura1001/mega_like_x/src/app/cli"
	"moura1001/mega_like_x/src/app/store"
	"os"
)

func main() {
	fmt.Println("Where's the floor?")
	fmt.Println("Type '{Game} like' to record a vote")

	poll, err := cli.NewCLI(store.IN_MEMORY, os.Stdin, nil)

	if err != nil {
		log.Fatalf("poll initialization error: %v", err)
	}

	poll.StartPoll()
}
