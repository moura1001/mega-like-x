package store

import (
	"moura1001/mega_like_x/src/app/model"
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("/games from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "x2", "Likes": 11},
			{"Name": "x3", "Likes": 10}
		]`)

		store := NewFileSystemGameStore(database)

		got := store.GetPolling()
		want := []model.Game{
			{Name: "x2", Likes: 11},
			{Name: "x3", Likes: 10},
		}

		AssertPolling(t, got, want)
	})
}
