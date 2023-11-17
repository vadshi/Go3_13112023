package main

import (
	"context"
	"flag"
	"fmt"
	"local/hw/db"
	"local/hw/handlers"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var seedFlag = flag.Bool("seed", false, "seed data to db")

func main() {
	flag.Parse()

	fmt.Println("Starting db connection...")
	pool, err := pgxpool.New(context.Background(), "user=postgres password=postgres dbname=db sslmode=disable host=localhost port=5435")
	if err != nil {
		log.Fatal("Can not connect to db", err)
	}
	defer pool.Close()

	store := db.NewStore(pool)

	if *seedFlag {
		db.Seed(store) // seed data to db
		return
	}

	server := handlers.NewServer(store)

	fmt.Println("Server is running... http://localhost:3000/")
	err = server.Start()
	if err != nil {
		log.Fatal("Can not start server", err)
	}

	// http.HandleFunc("/card-modal", handlerCardModal)
}

// func handlerCardModal(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Content-Type", "text/html")

// 	err := cardModal.Execute(w, nil)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
