package store

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
