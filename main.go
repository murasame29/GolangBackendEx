package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/murasame29/GolangBackendEx/api"
	db "github.com/murasame29/GolangBackendEx/db/sqlc"
)

const (
	DB_DRIVER      = "postgres"
	DB_SOURSE      = "postgresql://root:postgres@localhost:5432/simplebank?sslmode=disable"
	SERVER_ADDRESS = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(DB_DRIVER, DB_SOURSE)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NweStore(conn)
	server := api.NewServer(store)
	err = server.Start(SERVER_ADDRESS)
	if err != nil {
		log.Fatal("cannnot start server :", err)
	}
}
