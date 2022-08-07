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
	result, _ := p.DB.Exec("UPDATE games SET likes=(likes+1) WHERE name=$1", name)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected <= 0 {
		p.DB.Exec("INSERT INTO games(name, likes) VALUES($1, 1)", name)
	}
}

func (p *PostgresGameStore) GetPolling() model.Polling {
	rows, _ := p.DB.Query("SELECT name, likes FROM games")
	defer rows.Close()

	polling := model.Polling{}

	for rows.Next() {
		var game model.Game
		rows.Scan(&game.Name, &game.Likes)
		polling = append(polling, game)
	}

	return polling
}
