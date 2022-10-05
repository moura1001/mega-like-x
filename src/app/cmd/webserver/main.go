package main

import (
	"log"
	"moura1001/mega_like_x/src/app/alerter"
	"moura1001/mega_like_x/src/app/poll"
	"moura1001/mega_like_x/src/app/store"
	"moura1001/mega_like_x/src/app/webserver"
	"net/http"
)

func main() {
	store := store.NewInMemoryGameStore()
	poll := poll.NewMegaLike(
		store,
		alerter.BlindAlerterFunc(alerter.Alerter),
	)

	server, err := webserver.NewGameServer(store, "../../../templates/poll.html", poll)
	if err != nil {
		log.Fatalf("server startup error: %v", err)
	}

	if err := http.ListenAndServe(":4000", server); err != nil {
		log.Fatalf("could not listen on port 4000: %v", err)
	}
}
