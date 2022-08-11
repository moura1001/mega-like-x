package store_test

import (
	"moura1001/mega_like_x/src/app/model"
	"moura1001/mega_like_x/src/app/store"
	utilstestingfilestore "moura1001/mega_like_x/src/app/utils/test/file_store"
	utilstesting "moura1001/mega_like_x/src/app/utils/test/shared"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("/games from a reader", func(t *testing.T) {
		database, cleanDatabase := utilstestingfilestore.CreateTempFile(t, `[
			{"Name": "x2", "Likes": 11},
			{"Name": "x3", "Likes": 10}
		]`)
		defer cleanDatabase()

		store, err := store.NewFileSystemGameStore(database)
		utilstesting.AssertNoError(t, err)

		got := store.GetPolling()
		want := model.Polling{
			{Name: "x2", Likes: 11},
			{Name: "x3", Likes: 10},
		}

		utilstesting.AssertPolling(t, got, want)

		got = store.GetPolling()
		utilstesting.AssertPolling(t, got, want)
	})

	t.Run("get game likes", func(t *testing.T) {
		database, cleanDatabase := utilstestingfilestore.CreateTempFile(t, `[
			{"Name": "x7", "Likes": 3},
			{"Name": "x8", "Likes": 0}
		]`)
		defer cleanDatabase()

		store, err := store.NewFileSystemGameStore(database)
		utilstesting.AssertNoError(t, err)

		got := store.GetGameLikes("x7")
		want := 3

		utilstesting.AssertLikesValue(t, got, want)
	})

	t.Run("store likes for existing game", func(t *testing.T) {
		database, cleanDatabase := utilstestingfilestore.CreateTempFile(t, `[
			{"Name": "x1", "Likes": 6},
			{"Name": "x5", "Likes": 1}
		]`)
		defer cleanDatabase()

		store, err := store.NewFileSystemGameStore(database)
		utilstesting.AssertNoError(t, err)

		store.RecordLike("x1")

		got := store.GetGameLikes("x1")
		want := 7

		utilstesting.AssertLikesValue(t, got, want)
	})

	t.Run("store likes for new games", func(t *testing.T) {
		database, cleanDatabase := utilstestingfilestore.CreateTempFile(t, `[
			{"Name": "x4", "Likes": 0},
			{"Name": "x6", "Likes": 7}
		]`)
		defer cleanDatabase()

		store, err := store.NewFileSystemGameStore(database)
		utilstesting.AssertNoError(t, err)

		store.RecordLike("x1")

		got := store.GetGameLikes("x1")
		want := 1

		utilstesting.AssertLikesValue(t, got, want)
	})

	t.Run("polling sorted", func(t *testing.T) {
		database, cleanDatabase := utilstestingfilestore.CreateTempFile(t, `[
			{"Name": "corrupted", "Likes": 1},
			{"Name": "x7", "Likes": 11}
		]`)
		defer cleanDatabase()

		store, err := store.NewFileSystemGameStore(database)
		utilstesting.AssertNoError(t, err)

		got := store.GetPolling()
		want := model.Polling{
			{Name: "x7", Likes: 11},
			{Name: "corrupted", Likes: 1},
		}

		utilstesting.AssertPolling(t, got, want)

		// read again
		got = store.GetPolling()
		utilstesting.AssertPolling(t, got, want)
	})
}
