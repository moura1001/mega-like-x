package webserver_test

import (
	"moura1001/mega_like_x/src/app/model"
	utilstesting "moura1001/mega_like_x/src/app/utils/test/shared"
	"moura1001/mega_like_x/src/app/webserver"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETLikes(t *testing.T) {
	store := utilstesting.GetNewStubGameStore(
		map[string]int{
			"x1": 32,
			"x2": 64,
		},
		nil,
		nil,
	)
	server := webserver.NewGameServer(&store, "")

	t.Run("returns Mega Man X's likes", func(t *testing.T) {
		request := utilstesting.NewGetLikesRequest("x1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utilstesting.AssertStatus(t, response.Code, http.StatusOK)
		utilstesting.AssertResponseBody(t, response.Body.String(), "32")
	})

	t.Run("returns Mega Man X2's likes", func(t *testing.T) {
		request := utilstesting.NewGetLikesRequest("x2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utilstesting.AssertStatus(t, response.Code, http.StatusOK)
		utilstesting.AssertResponseBody(t, response.Body.String(), "64")
	})

	t.Run("returns 404 on missing game", func(t *testing.T) {
		request := utilstesting.NewGetLikesRequest("corrupted")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utilstesting.AssertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreLikes(t *testing.T) {
	store := utilstesting.GetNewStubGameStore(
		map[string]int{},
		nil,
		nil,
	)
	server := webserver.NewGameServer(&store, "")

	t.Run("it records likes when POST", func(t *testing.T) {
		game := "x6"

		request := utilstesting.NewPostLikeRequest(game)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utilstesting.AssertStatus(t, response.Code, http.StatusAccepted)

		utilstesting.AssertGameLike(t, &store, "x6")

	})
}

func TestPolling(t *testing.T) {
	wantedGames := model.Polling{
		{Name: "x1", Likes: 30},
		{Name: "x4", Likes: 12},
		{Name: "x6", Likes: 23},
	}

	store := utilstesting.GetNewStubGameStore(nil, nil, wantedGames)
	server := webserver.NewGameServer(&store, "")

	t.Run("it returns the game table as JSON", func(t *testing.T) {

		request := utilstesting.NewGetPollingRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utilstesting.AssertContentType(t, response, utilstesting.JsonContentType)

		got := utilstesting.GetPollingFromResponse(t, response.Body)

		utilstesting.AssertStatus(t, response.Code, http.StatusOK)
		utilstesting.AssertPolling(t, got, wantedGames)
	})
}

func TestPoll(t *testing.T) {
	store := utilstesting.GetNewStubGameStore(nil, nil, model.Polling{})
	server := webserver.NewGameServer(&store, "../../templates/poll.html")

	t.Run("GET /poll returns 200", func(t *testing.T) {

		request := utilstesting.NewPollRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		utilstesting.AssertStatus(t, response.Code, http.StatusOK)
	})
}
