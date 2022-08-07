package store

import (
	"io"
	"io/ioutil"
	"moura1001/mega_like_x/src/app/model"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("/games from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "x2", "Likes": 11},
			{"Name": "x3", "Likes": 10}
		]`)
		defer cleanDatabase()

		store := NewFileSystemGameStore(database)

		got := store.GetPolling()
		want := []model.Game{
			{Name: "x2", Likes: 11},
			{Name: "x3", Likes: 10},
		}

		AssertPolling(t, got, want)

		got = store.GetPolling()
		AssertPolling(t, got, want)
	})

	t.Run("get game likes", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "x7", "Likes": 3},
			{"Name": "x8", "Likes": 0}
		]`)
		defer cleanDatabase()

		store := NewFileSystemGameStore(database)

		got := store.GetGameLikes("x7")
		want := 3

		AssertLikesValue(t, got, want)
	})
}

func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
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
