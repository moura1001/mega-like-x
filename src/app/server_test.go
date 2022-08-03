package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETLikes(t *testing.T) {
	t.Run("returns Mega Man X's likes", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/likes/x1", nil)
		response := httptest.NewRecorder()

		Server(response, request)

		got := response.Body.String()
		want := "32"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns Mega Man X2's likes", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/likes/x2", nil)
		response := httptest.NewRecorder()

		Server(response, request)

		got := response.Body.String()
		want := "64"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}
