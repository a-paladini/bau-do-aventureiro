package main

import (
	"database/sql"
	"log"

	"github.com/a-paladini/bau-do-aventureiro/api"
	db "github.com/a-paladini/bau-do-aventureiro/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/bau_t20?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
