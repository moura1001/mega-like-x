package store

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"moura1001/mega_like_x/src/app/model"
	"os"
	"reflect"
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

func AssertPolling(t *testing.T, got, want model.Polling) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertLikesValue(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct likes, got %d, want %d", got, want)
	}
}

func CreateTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("didnt expect an error but got one: %v", err)
	}
}
