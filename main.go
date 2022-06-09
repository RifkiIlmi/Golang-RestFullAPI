package main

import (
	"database/sql"
	"log"

	"github.com/Rifkiilmi/simplebank/api"
	db "github.com/Rifkiilmi/simplebank/db/sqlc"
	"github.com/Rifkiilmi/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Run(config.ServerAddr)
	if err != nil {
		log.Fatal("cannot run server", err)
	}
}
