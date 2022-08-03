package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubGameStore struct {
	likes map[string]int
}

func (s *StubGameStore) GetGameLikes(name string) int {
	return s.likes[name]
}

func TestGETLikes(t *testing.T) {
	store := StubGameStore{
		map[string]int{
			"x1": 32,
			"x2": 64,
		},
	}
	server := &GameServer{&store}

	t.Run("returns Mega Man X's likes", func(t *testing.T) {
		request := newGetLikesRequest("x1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "32")
	})

	t.Run("returns Mega Man X2's likes", func(t *testing.T) {
		request := newGetLikesRequest("x2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "64")
	})
}

func newGetLikesRequest(game string) *http.Request {
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/likes/%s", game), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body error, got '%s', want '%s'", got, want)
	}
}
