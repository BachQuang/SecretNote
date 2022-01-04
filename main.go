package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/secretnote/backend/api"
	db "github.com/secretnote/backend/db/sqlc"
	"github.com/secretnote/backend/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not load the config", err)
	}

	conn, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.SERVER_ADDRESS)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
