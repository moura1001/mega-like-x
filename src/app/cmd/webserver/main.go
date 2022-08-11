package main

import (
	"log"
	"moura1001/mega_like_x/src/app/store"
	"moura1001/mega_like_x/src/app/webserver"
	"net/http"
)

func main() {
	server, err := webserver.NewGameServer(store.IN_MEMORY, nil)

	if err != nil {
		log.Fatalf("server initialization error: %v", err)
	}

	if err := http.ListenAndServe(":4000", server); err != nil {
		log.Fatalf("could not listen on port 4000: %v", err)
	}
}
