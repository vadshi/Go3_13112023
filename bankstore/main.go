package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vadshi/go3/bankstore/api"
	db "github.com/vadshi/go3/bankstore/db/sqlc"
)

const (
	dbSource      = "postgresql://postgres:postgres@localhost:5435/bankstoredb?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	pool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Can not connect to db", err)
	}

	defer pool.Close()

	store := db.NewStore(pool)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Can not start server", err)
	}

}
