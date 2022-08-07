package app

import (
	"moura1001/mega_like_x/src/app/model"
	"moura1001/mega_like_x/src/app/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingLikesAndRetrievingThemMemory(t *testing.T) {
	server := NewGameServer(store.IN_MEMORY)
	game := "x4"

	server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(game))
	server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(game))
	server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(game))
	server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(game))

	t.Run("get likes", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetLikesRequest(game))

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "4")
	})

	t.Run("get polling", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetPollingRequest())

		assertStatus(t, response.Code, http.StatusOK)
		got := getPollingFromResponse(t, response.Body)
		want := model.Polling{
			{Name: "x4", Likes: 4},
		}
		store.AssertPolling(t, got, want)
	})
}

func TestRecordingLikesAndRetrievingThemFromPostgres(t *testing.T) {
	server := NewGameServer(store.POSTGRES)
	server.store = store.SetupPostgresStoreTests(t)

	game := "x8"
	// cast to Postgres store
	st := server.store.(*store.PostgresGameStore)
	st.DB.Exec(`
		INSERT INTO	games(name)
		VALUES	($1)
	`, game)

	server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(game))
	server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(game))

	t.Run("get likes", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetLikesRequest(game))

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "2")
	})

	t.Run("get polling", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetPollingRequest())

		assertStatus(t, response.Code, http.StatusOK)
		got := getPollingFromResponse(t, response.Body)
		want := model.Polling{
			{Name: "x8", Likes: 2},
		}
		store.AssertPolling(t, got, want)
	})
}
