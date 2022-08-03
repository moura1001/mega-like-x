package app

import (
	"fmt"
	"net/http"
)

func Server(w http.ResponseWriter, r *http.Request) {
	game := r.URL.Path[len("/likes/"):]

	fmt.Fprint(w, GetGameLikes(game))
}

func GetGameLikes(name string) string {
	if name == "x1" {
		return "32"
	}

	if name == "x2" {
		return "64"
	}

	return ""
}
