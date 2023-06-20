package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	DB_DRIVER = "postgres"
	DB_SOURSE = "postgresql://root:postgres@localhost:5432/simplebank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(DB_DRIVER, DB_SOURSE)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
