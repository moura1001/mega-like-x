package store

import "database/sql"

type PostgresGameStore struct {
	DB *sql.DB
}

func NewPostgresGameStore() *PostgresGameStore {
	return &PostgresGameStore{}
}

func (p *PostgresGameStore) GetGameLikes(name string) int {
	var likes int
	p.DB.QueryRow("SELECT likes FROM games WHERE name=$1", name).Scan(&likes)
	return likes
}

func (p *PostgresGameStore) RecordLike(name string) {}
