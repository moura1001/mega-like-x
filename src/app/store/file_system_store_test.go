package store

import (
	"moura1001/mega_like_x/src/app/model"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("/games from a reader", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "x2", "Likes": 11},
			{"Name": "x3", "Likes": 10}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemGameStore(database)
		AssertNoError(t, err)

		got := store.GetPolling()
		want := model.Polling{
			{Name: "x2", Likes: 11},
			{Name: "x3", Likes: 10},
		}

		AssertPolling(t, got, want)

		got = store.GetPolling()
		AssertPolling(t, got, want)
	})

	t.Run("get game likes", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "x7", "Likes": 3},
			{"Name": "x8", "Likes": 0}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemGameStore(database)
		AssertNoError(t, err)

		got := store.GetGameLikes("x7")
		want := 3

		AssertLikesValue(t, got, want)
	})

	t.Run("store likes for existing game", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "x1", "Likes": 6},
			{"Name": "x5", "Likes": 1}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemGameStore(database)
		AssertNoError(t, err)

		store.RecordLike("x1")

		got := store.GetGameLikes("x1")
		want := 7

		AssertLikesValue(t, got, want)
	})

	t.Run("store likes for new games", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "x4", "Likes": 0},
			{"Name": "x6", "Likes": 7}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemGameStore(database)
		AssertNoError(t, err)

		store.RecordLike("x1")

		got := store.GetGameLikes("x1")
		want := 1

		AssertLikesValue(t, got, want)
	})

	t.Run("polling sorted", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "corrupted", "Likes": 1},
			{"Name": "x7", "Likes": 11}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemGameStore(database)
		AssertNoError(t, err)

		got := store.GetPolling()
		want := model.Polling{
			{Name: "x7", Likes: 11},
			{Name: "corrupted", Likes: 1},
		}

		AssertPolling(t, got, want)

		// read again
		got = store.GetPolling()
		AssertPolling(t, got, want)
	})
}
