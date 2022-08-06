package store

import "moura1001/mega_like_x/src/app/model"

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

func (i *InMemoryGameStore) GetPolling() []model.Game {
	var polling []model.Game
	for name, likes := range i.store {
		polling = append(polling, model.Game{Name: name, Likes: likes})
	}
	return polling
}
