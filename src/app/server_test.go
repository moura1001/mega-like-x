package app

import (
	"encoding/json"
	"fmt"
	"io"
	"moura1001/mega_like_x/src/app/model"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubGameStore struct {
	likes     map[string]int
	likeCalls []string
	polling   []model.Game
}

func (s *StubGameStore) GetGameLikes(name string) int {
	return s.likes[name]
}

func (s *StubGameStore) RecordLike(name string) {
	s.likeCalls = append(s.likeCalls, name)
}

func (s *StubGameStore) GetPolling() []model.Game {
	return s.polling
}

func TestGETLikes(t *testing.T) {
	store := StubGameStore{
		map[string]int{
			"x1": 32,
			"x2": 64,
		},
		nil,
		nil,
	}
	server := NewGameServer("")
	server.store = &store

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
		nil,
	}
	server := NewGameServer("")
	server.store = &store

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

func TestPolling(t *testing.T) {
	wantedGames := []model.Game{
		{Name: "x1", Likes: 30},
		{Name: "x4", Likes: 12},
		{Name: "x6", Likes: 23},
	}

	store := StubGameStore{nil, nil, wantedGames}
	server := NewGameServer("")
	server.store = &store

	t.Run("it returns the game table as JSON", func(t *testing.T) {

		request := newGetPollingRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertContentType(t, response, jsonContentType)

		got := getPollingFromResponse(t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)
		assertPolling(t, got, wantedGames)
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

func newGetPollingRequest() *http.Request {
	req := httptest.NewRequest(http.MethodGet, "/games", nil)
	return req
}

func getPollingFromResponse(t *testing.T, body io.Reader) (polling []model.Game) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&polling)
	if err != nil {
		t.Fatalf("unable to parse response from server '%v' into slice of Vote: '%v'", body, err)
	}

	return
}

func assertPolling(t *testing.T, got, want []model.Game) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

const jsonContentType = "application/json"

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Body)
	}
}
