package store

type PostgresGameStore struct{}

func NewPostgresGameStore() *PostgresGameStore {
	return &PostgresGameStore{}
}

func (p *PostgresGameStore) GetGameLikes(name string) int {
	return 7
}

func (p *PostgresGameStore) RecordLike(name string) {}
