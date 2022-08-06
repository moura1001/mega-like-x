package store

import "testing"

func TestConnectionPing(t *testing.T) {
	store := SetupPostgresStoreTests(t)

	err := store.DB.Ping()
	if err != nil {
		t.Errorf("postgres ping connection error: '%v'", err)
	}
}

func TestGETLikes(t *testing.T) {
	store := SetupPostgresStoreTests(t)

	store.DB.Exec(`
		INSERT INTO
			games(name, likes)
		VALUES
			('x3', 8),
			('x7', 1)
	`)

	t.Run("returns Mega Man X3's likes", func(t *testing.T) {
		likes := store.GetGameLikes("x3")

		assertLikesValue(t, likes, 8)
	})

	t.Run("returns Mega Man X7's likes", func(t *testing.T) {
		likes := store.GetGameLikes("x7")

		assertLikesValue(t, likes, 1)
	})
}

func TestStoreLikes(t *testing.T) {
	store := SetupPostgresStoreTests(t)

	store.DB.Exec(`
		INSERT INTO	games(name)
		VALUES	('x4')
	`)

	t.Run("record user like", func(t *testing.T) {
		game := "x4"

		likes := store.GetGameLikes(game)
		assertLikesValue(t, likes, 0)

		store.RecordLike(game)

		likes = store.GetGameLikes(game)
		assertLikesValue(t, likes, 1)
	})
}

func assertLikesValue(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct likes, got %d, want %d", got, want)
	}
}
