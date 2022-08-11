package webserver

import (
	"fmt"
	"io"
	"moura1001/mega_like_x/src/app/model"
	"moura1001/mega_like_x/src/app/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETLikes(t *testing.T) {
	st := store.StubGameStore{
		Likes: map[string]int{
			"x1": 32,
			"x2": 64,
		},
		LikeCalls: nil,
		Polling:   nil,
	}
	server, _ := NewGameServer("", nil)
	server.store = &st

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
	st := store.StubGameStore{
		Likes:     map[string]int{},
		LikeCalls: nil,
		Polling:   nil,
	}
	server, _ := NewGameServer("", nil)
	server.store = &st

	t.Run("it records likes when POST", func(t *testing.T) {
		game := "x6"

		request := newPostLikeRequest(game)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		store.AssertGameLike(t, &st, "x6")

	})
}

func TestPolling(t *testing.T) {
	wantedGames := model.Polling{
		{Name: "x1", Likes: 30},
		{Name: "x4", Likes: 12},
		{Name: "x6", Likes: 23},
	}

	st := store.StubGameStore{Likes: nil, LikeCalls: nil, Polling: wantedGames}
	server, _ := NewGameServer("", nil)
	server.store = &st

	t.Run("it returns the game table as JSON", func(t *testing.T) {

		request := newGetPollingRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertContentType(t, response, jsonContentType)

		got := getPollingFromResponse(t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)
		store.AssertPolling(t, got, wantedGames)
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

func getPollingFromResponse(t *testing.T, body io.Reader) (polling model.Polling) {
	t.Helper()

	polling, err := model.NewGamePolling(body)
	if err != nil {
		t.Fatalf("unable to parse response from server '%v' to get polling: '%v'", body, err)
	}

	return
}

const jsonContentType = "application/json"

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Body)
	}
}
