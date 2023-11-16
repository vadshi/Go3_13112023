/*
Full database
1. Run go run . -seed
Run app
2. Run go run .
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"local/hw1/db"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

var layout = template.Must(template.ParseFiles(
	"layout.html.tmpl",
	"card.html.tmpl",
	"card_empty.html.tmpl",
))

var seedFlag = flag.Bool("seed", false, "seed data to db")
var ctx = context.Background()
var queries *db.Queries

func main() {

	fmt.Println("Starting db connection...")
	conn, err := pgx.Connect(ctx, "user=postgres password=postgres dbname=db sslmode=disable host=localhost port=5435")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	queries = db.New(conn)

	flag.Parse()
	if *seedFlag {
		seed() // seed data to db
		return
	}

	fmt.Println("Server is running... http://localhost:3000/")
	http.HandleFunc("/", handlerCardsList)
	http.HandleFunc("/new", handlerCardNew)
	http.ListenAndServe(":3000", nil)
}

func handlerCardsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	cards, err := queries.ListCards(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = layout.Execute(w, cards)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerCardNew(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// TODO
	w.Write([]byte("OK"))
}
