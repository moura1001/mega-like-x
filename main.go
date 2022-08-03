package main

import (
	"log"
	"moura1001/mega_like_x/src/app"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(app.Server)
	if err := http.ListenAndServe(":4000", handler); err != nil {
		log.Fatalf("could not listen on port 4000: %v", err)
	}
}
