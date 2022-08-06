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

func TestRecordingLikesAndRetrievingThemFromPostgres(t *testing.T) {
	server := NewGameServer(store.POSTGRES)
	server.store = store.SetupPostgresStoreTests(t)

	game := "x8"
	// cast to Postgres store
	store := server.store.(*store.PostgresGameStore)
	store.DB.Exec(`
		INSERT INTO	games(name)
		VALUES	($1)
	`, game)

	server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(game))
	server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(game))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetLikesRequest(game))

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "2")
}
