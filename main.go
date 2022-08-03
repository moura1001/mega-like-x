package main

import (
	"log"
	"moura1001/mega_like_x/src/app"
	"net/http"
)

func main() {
	server := &app.GameServer{}

	if err := http.ListenAndServe(":4000", server); err != nil {
		log.Fatalf("could not listen on port 4000: %v", err)
	}
}
