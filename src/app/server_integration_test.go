package app

import (
	"moura1001/mega_like_x/src/app/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingLikesAndRetrievingThem(t *testing.T) {
	server := NewGameServer(store.IN_MEMORY)
	game := "x4"

	server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(game))
	server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(game))
	server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(game))
	server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(game))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetLikesRequest(game))

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "4")
}
