package store

type InMemoryGameStore struct{}

func (i *InMemoryGameStore) GetGameLikes(name string) int {
	return 7
}

func (i *InMemoryGameStore) RecordLike(name string) {}
