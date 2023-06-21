package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/murasame29/GolangBackendEx/api"
	db "github.com/murasame29/GolangBackendEx/db/sqlc"
	"github.com/murasame29/GolangBackendEx/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannnot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSouse)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NweStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannnot start server :", err)
	}
}
