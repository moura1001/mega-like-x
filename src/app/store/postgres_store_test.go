package store_test

import (
	"moura1001/mega_like_x/src/app/model"
	utilstestingpgstore "moura1001/mega_like_x/src/app/utils/test/pg_store"
	utilstesting "moura1001/mega_like_x/src/app/utils/test/shared"
	"testing"
)

func TestConnectionPing(t *testing.T) {
	store := utilstestingpgstore.SetupPostgresStoreTests(t)

	err := store.DB.Ping()
	if err != nil {
		t.Errorf("postgres ping connection error: '%v'", err)
	}
}

func TestGETLikes(t *testing.T) {
	store := utilstestingpgstore.SetupPostgresStoreTests(t)

	store.DB.Exec(`
		INSERT INTO
			games(name, likes)
		VALUES
			('x3', 8),
			('x7', 1)
	`)

	t.Run("returns Mega Man X3's likes", func(t *testing.T) {
		likes := store.GetGameLikes("x3")

		utilstesting.AssertLikesValue(t, likes, 8)
	})

	t.Run("returns Mega Man X7's likes", func(t *testing.T) {
		likes := store.GetGameLikes("x7")

		utilstesting.AssertLikesValue(t, likes, 1)
	})
}

func TestStoreLikes(t *testing.T) {
	store := utilstestingpgstore.SetupPostgresStoreTests(t)

	store.DB.Exec(`
		INSERT INTO	games(name)
		VALUES	('x4')
	`)

	t.Run("record user like", func(t *testing.T) {
		game := "x4"

		likes := store.GetGameLikes(game)
		utilstesting.AssertLikesValue(t, likes, 0)

		store.RecordLike(game)

		likes = store.GetGameLikes(game)
		utilstesting.AssertLikesValue(t, likes, 1)
	})

	t.Run("record like to new games", func(t *testing.T) {
		game := "x7"

		likes := store.GetGameLikes(game)
		utilstesting.AssertLikesValue(t, likes, 0)

		store.RecordLike(game)

		likes = store.GetGameLikes(game)
		utilstesting.AssertLikesValue(t, likes, 1)
	})

	t.Run("likes sorted", func(t *testing.T) {
		store.RecordLike("x7")
		store.RecordLike("x7")

		got := store.GetPolling()
		want := model.Polling{
			{Name: "x7", Likes: 3},
			{Name: "x4", Likes: 1},
		}

		utilstesting.AssertPolling(t, got, want)
	})
}
