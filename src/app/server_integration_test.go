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
	newGame := "x6"

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

	t.Run("record user like to new games", func(t *testing.T) {
		response := httptest.NewRecorder()

		server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(newGame))
		server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(newGame))

		server.ServeHTTP(response, newGetPollingRequest())

		assertStatus(t, response.Code, http.StatusOK)
		got := getPollingFromResponse(t, response.Body)
		want := model.Polling{
			{Name: "x4", Likes: 4},
			{Name: "x6", Likes: 2},
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

func TestRecordingLikesAndRetrievingThemFromFile(t *testing.T) {
	database, cleanDatabase := store.CreateTempFile(t, `[
		{"Name": "x1", "Likes": 4}
	]`)
	defer cleanDatabase()

	st := store.NewFileSystemGameStore(database)
	server := NewGameServer(store.FILE_SYSTEM)
	server.store = st

	game := "x1"
	newGame := "x2"

	server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(game))
	server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(game))

	t.Run("get likes", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetLikesRequest(game))

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "6")
	})

	t.Run("get polling", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetPollingRequest())

		assertStatus(t, response.Code, http.StatusOK)
		got := getPollingFromResponse(t, response.Body)
		want := model.Polling{
			{Name: "x1", Likes: 6},
		}
		store.AssertPolling(t, got, want)
	})

	t.Run("record user like to new games", func(t *testing.T) {
		response := httptest.NewRecorder()

		server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(newGame))
		server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(newGame))
		server.ServeHTTP(httptest.NewRecorder(), newPostLikeRequest(newGame))

		server.ServeHTTP(response, newGetPollingRequest())

		assertStatus(t, response.Code, http.StatusOK)
		got := getPollingFromResponse(t, response.Body)
		want := model.Polling{
			{Name: "x1", Likes: 6},
			{Name: "x2", Likes: 3},
		}
		store.AssertPolling(t, got, want)
	})
}
