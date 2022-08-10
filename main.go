package main

import (
	"log"
	"moura1001/mega_like_x/src/app"
	"moura1001/mega_like_x/src/app/store"
	"net/http"
)

func main() {
	const storeType = store.IN_MEMORY
	server, err := app.NewGameServer(storeType, nil)

	if err != nil {
		log.Fatalf("problem creating '%s' game store: %v", storeType, err)
	}

	if err := http.ListenAndServe(":4000", server); err != nil {
		log.Fatalf("could not listen on port 4000: %v", err)
	}
}
