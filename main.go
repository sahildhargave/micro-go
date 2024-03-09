package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/sahil/simplebank/api"
	db "github.com/sahil/simplebank/db/sqlc"
	"github.com/sahil/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil{
	log.Fatal("cannot connect to DB:", err)
	}
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
