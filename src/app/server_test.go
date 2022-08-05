package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubGameStore struct {
	likes     map[string]int
	likeCalls []string
}

func (s *StubGameStore) GetGameLikes(name string) int {
	return s.likes[name]
}

func (s *StubGameStore) RecordLike(name string) {
	s.likeCalls = append(s.likeCalls, name)
}

func TestGETLikes(t *testing.T) {
	store := StubGameStore{
		map[string]int{
			"x1": 32,
			"x2": 64,
		},
		nil,
	}
	server := &GameServer{&store}

	t.Run("returns Mega Man X's likes", func(t *testing.T) {
		request := newGetLikesRequest("x1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "32")
	})

	t.Run("returns Mega Man X2's likes", func(t *testing.T) {
		request := newGetLikesRequest("x2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "64")
	})

	t.Run("returns 404 on missing game", func(t *testing.T) {
		request := newGetLikesRequest("corrupted")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreLikes(t *testing.T) {
	store := StubGameStore{
		map[string]int{},
		nil,
	}
	server := &GameServer{&store}

	t.Run("it records likes when POST", func(t *testing.T) {
		game := "x6"

		request := newPostLikeRequest(game)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.likeCalls) != 1 {
			t.Errorf("got %d calls to RecordLike, want %d", len(store.likeCalls), 1)
		}

		if store.likeCalls[0] != game {
			t.Errorf("did not store correct liked game, got '%s' want '%s'", store.likeCalls[0], game)
		}
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

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got '%d', want '%d'", got, want)
	}
}

func newPostLikeRequest(game string) *http.Request {
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/likes/%s", game), nil)
	return req
}
