package store

import (
	"moura1001/mega_like_x/src/app/model"
	"sort"
)

type InMemoryGameStore struct {
	store map[string]int
}

func NewInMemoryGameStore() *InMemoryGameStore {
	return &InMemoryGameStore{map[string]int{}}
}

func (i *InMemoryGameStore) GetGameLikes(name string) int {
	return i.store[name]
}

func (i *InMemoryGameStore) RecordLike(name string) {
	i.store[name]++
}

func (i *InMemoryGameStore) GetPolling() model.Polling {
	var polling model.Polling
	for name, likes := range i.store {
		polling = append(polling, model.Game{Name: name, Likes: likes})
	}

	sort.Slice(polling, func(i, j int) bool {
		return polling[i].Likes > polling[j].Likes
	})
	return polling
}
