package store

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	user     = "usertest"
	password = "test"
	dbname   = "dbtest"
)

var store *PostgresGameStore

func setup(t *testing.T) {
	store = NewPostgresGameStore()
	store.DB = getPostgresConnection(t)

	store.DB.Exec(`
		CREATE TABLE IF NOT EXISTS games(
			name TEXT NOT NULL,
			likes BIGINT NOT NULL DEFAULT 0,
			CONSTRAINT games_pkey PRIMARY KEY (name)
		)
	`)

	store.DB.Exec("TRUNCATE games")
}

func TestConnectionPing(t *testing.T) {
	setup(t)

	err := store.DB.Ping()
	if err != nil {
		t.Errorf("postgres ping connection error: '%v'", err)
	}
}

func TestGETLikes(t *testing.T) {
	setup(t)

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
	setup(t)

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

func getPostgresConnection(t *testing.T) *sql.DB {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s "+
		"dbname=%s sslmode=disable", host, user, password, dbname)

	pgConn, err := sql.Open("postgres", connectionString)
	if err != nil {
		t.Fatalf("get connection error: '%v'", err)
	}

	return pgConn
}

func assertLikesValue(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct likes, got %d, want %d", got, want)
	}
}
