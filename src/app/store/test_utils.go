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

func SetupPostgresStoreTests(t *testing.T) *PostgresGameStore {
	t.Helper()

	store := NewPostgresGameStore()
	store.DB = getPostgresConnection(t)

	store.DB.Exec(`
		CREATE TABLE IF NOT EXISTS games(
			name TEXT NOT NULL,
			likes BIGINT NOT NULL DEFAULT 0,
			CONSTRAINT games_pkey PRIMARY KEY (name)
		)
	`)

	store.DB.Exec("TRUNCATE games")

	return store
}

func getPostgresConnection(t *testing.T) *sql.DB {
	t.Helper()

	connectionString := fmt.Sprintf("host=%s user=%s password=%s "+
		"dbname=%s sslmode=disable", host, user, password, dbname)

	pgConn, err := sql.Open("postgres", connectionString)
	if err != nil {
		t.Fatalf("get connection error: '%v'", err)
	}

	return pgConn
}
