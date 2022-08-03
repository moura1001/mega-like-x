package app

import (
	"fmt"
	"net/http"
)

func Server(w http.ResponseWriter, r *http.Request) {
	game := r.URL.Path[len("/likes/"):]

	if game == "x1" {
		fmt.Fprint(w, "32")
		return
	}

	if game == "x2" {
		fmt.Fprint(w, "64")
		return
	}

}
