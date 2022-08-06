package store

import (
	"database/sql"
	"moura1001/mega_like_x/src/app/model"
)

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

func (p *PostgresGameStore) RecordLike(name string) {
	p.DB.Exec("UPDATE games SET likes=(likes+1) WHERE name=$1", name)
}

func (i *PostgresGameStore) GetPolling() []model.Game {
	return nil
}
