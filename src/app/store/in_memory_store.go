package store

type InMemoryGameStore struct{}

func (i *InMemoryGameStore) GetGameLikes(name string) int {
	return 7
}
