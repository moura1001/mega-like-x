package webserver

import (
	"moura1001/mega_like_x/src/app/model"
	"moura1001/mega_like_x/src/app/store"
	utilstestingfilestore "moura1001/mega_like_x/src/app/utils/test/file_store"
	utilstestingpgstore "moura1001/mega_like_x/src/app/utils/test/pg_store"
	utilstesting "moura1001/mega_like_x/src/app/utils/test/shared"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingLikesAndRetrievingThemMemory(t *testing.T) {
	server, err := NewGameServer(store.IN_MEMORY, nil)
	utilstesting.AssertNoError(t, err)

	game := "x4"
	newGame := "x6"

	server.ServeHTTP(httptest.NewRecorder(), utilstesting.NewPostLikeRequest(game))
	server.ServeHTTP(httptest.NewRecorder(), utilstesting.NewPostLikeRequest(game))
	server.ServeHTTP(httptest.NewRecorder(), utilstesting.NewPostLikeRequest(game))
	server.ServeHTTP(httptest.NewRecorder(), utilstesting.NewPostLikeRequest(game))

	t.Run("get likes", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, utilstesting.NewGetLikesRequest(game))

		utilstesting.AssertStatus(t, response.Code, http.StatusOK)
		utilstesting.AssertResponseBody(t, response.Body.String(), "4")
	})

	t.Run("get polling", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, utilstesting.NewGetPollingRequest())

		utilstesting.AssertStatus(t, response.Code, http.StatusOK)
		got := utilstesting.GetPollingFromResponse(t, response.Body)
		want := model.Polling{
			{Name: "x4", Likes: 4},
		}
		utilstesting.AssertPolling(t, got, want)
	})

	t.Run("record user like to new games", func(t *testing.T) {
		response := httptest.NewRecorder()

		server.ServeHTTP(httptest.NewRecorder(), utilstesting.NewPostLikeRequest(newGame))
		server.ServeHTTP(httptest.NewRecorder(), utilstesting.NewPostLikeRequest(newGame))

		server.ServeHTTP(response, utilstesting.NewGetPollingRequest())

		utilstesting.AssertStatus(t, response.Code, http.StatusOK)
		got := utilstesting.GetPollingFromResponse(t, response.Body)
		want := model.Polling{
			{Name: "x4", Likes: 4},
			{Name: "x6", Likes: 2},
		}
		utilstesting.AssertPolling(t, got, want)
	})
}

func TestRecordingLikesAndRetrievingThemFromPostgres(t *testing.T) {
	server, err := NewGameServer(store.POSTGRES, nil)
	utilstesting.AssertNoError(t, err)

	server.store = utilstestingpgstore.SetupPostgresStoreTests(t)

	game := "x8"
	// cast to Postgres store
	st := server.store.(*store.PostgresGameStore)
	st.DB.Exec(`
		INSERT INTO	games(name)
		VALUES	($1)
	`, game)

	server.ServeHTTP(httptest.NewRecorder(), utilstesting.NewPostLikeRequest(game))
	server.ServeHTTP(httptest.NewRecorder(), utilstesting.NewPostLikeRequest(game))

	t.Run("get likes", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, utilstesting.NewGetLikesRequest(game))

		utilstesting.AssertStatus(t, response.Code, http.StatusOK)
		utilstesting.AssertResponseBody(t, response.Body.String(), "2")
	})

	t.Run("get polling", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, utilstesting.NewGetPollingRequest())

		utilstesting.AssertStatus(t, response.Code, http.StatusOK)
		got := utilstesting.GetPollingFromResponse(t, response.Body)
		want := model.Polling{
			{Name: "x8", Likes: 2},
		}
		utilstesting.AssertPolling(t, got, want)
	})
}

func TestRecordingLikesAndRetrievingThemFromFile(t *testing.T) {
	database, cleanDatabase := utilstestingfilestore.CreateTempFile(t, `[
		{"Name": "x1", "Likes": 4}
	]`)
	defer cleanDatabase()

	server, err := NewGameServer(store.FILE_SYSTEM, database)
	utilstesting.AssertNoError(t, err)

	game := "x1"
	newGame := "x2"

	server.ServeHTTP(httptest.NewRecorder(), utilstesting.NewPostLikeRequest(game))
	server.ServeHTTP(httptest.NewRecorder(), utilstesting.NewPostLikeRequest(game))

	t.Run("get likes", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, utilstesting.NewGetLikesRequest(game))

		utilstesting.AssertStatus(t, response.Code, http.StatusOK)
		utilstesting.AssertResponseBody(t, response.Body.String(), "6")
	})

	t.Run("get polling", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, utilstesting.NewGetPollingRequest())

		utilstesting.AssertStatus(t, response.Code, http.StatusOK)
		got := utilstesting.GetPollingFromResponse(t, response.Body)
		want := model.Polling{
			{Name: "x1", Likes: 6},
		}
		utilstesting.AssertPolling(t, got, want)
	})

	t.Run("record user like to new games", func(t *testing.T) {
		response := httptest.NewRecorder()

		server.ServeHTTP(httptest.NewRecorder(), utilstesting.NewPostLikeRequest(newGame))
		server.ServeHTTP(httptest.NewRecorder(), utilstesting.NewPostLikeRequest(newGame))
		server.ServeHTTP(httptest.NewRecorder(), utilstesting.NewPostLikeRequest(newGame))

		server.ServeHTTP(response, utilstesting.NewGetPollingRequest())

		utilstesting.AssertStatus(t, response.Code, http.StatusOK)
		got := utilstesting.GetPollingFromResponse(t, response.Body)
		want := model.Polling{
			{Name: "x1", Likes: 6},
			{Name: "x2", Likes: 3},
		}
		utilstesting.AssertPolling(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := utilstestingfilestore.CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := store.NewFileSystemGameStore(database)

		utilstesting.AssertNoError(t, err)
	})
}
