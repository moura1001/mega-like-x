package store

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	user     = "usertest"
	password = "test"
	dbname   = "dbtest"
)

func TestConnectionPing(t *testing.T) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s "+
		"dbname=%s sslmode=disable", host, user, password, dbname)

	pgConn, err := sql.Open("postgres", connectionString)
	if err != nil {
		t.Fatalf("get connection error: '%v'", err)
	}
	defer pgConn.Close()

	err = pgConn.Ping()
	if err != nil {
		t.Errorf("postgres ping connection error: '%v'", err)
	}
}
