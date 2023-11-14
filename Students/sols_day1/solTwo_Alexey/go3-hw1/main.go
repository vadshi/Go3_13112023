package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

var layout = template.Must(template.ParseFiles(
	"layout.html.tmpl",
	"card.html.tmpl",
	"card_empty.html.tmpl",
))

var times = []string{
	"08:00",
	"09:00",
	"10:00",
	"11:00",
}

var dates = []string{
	"2023-11-13",
	"2023-11-14",
	"2023-11-15",
}

var users = []string{
	"women/1.jpg",
	"women/2.jpg",
	"women/3.jpg",
	"women/4.jpg",
	"men/1.jpg",
	"men/2.jpg",
	"men/3.jpg",
	"men/4.jpg",
}

var places = []string{
	"SM",
	"MC",
	"NA",
	"NC",
	"SM",
	"VK",
	"TGG",
	"TGP",
}

type Card struct {
	Time   string
	Date   string
	Title  string
	User   string
	Places []string
}

func main() {

	http.HandleFunc("/", handler)
	fmt.Println("Server is running... http://localhost:3000/")
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	//	placesRand := []string{}
	//	for _, place := range places {
	//		if rand.Intn(2) == 1 {
	//			placesRand = append(placesRand, place)
	//		}
	//	}

	//	if len(placesRand) > 4 {
	//		placesRand = placesRand[:4]
	//	}

	cards := []Card{}
	cardsCount := rand.Intn(3) + 1
	for i := 0; i < cardsCount; i++ {
		card := Card{
			Time:   times[rand.Intn(len(times))],
			Date:   dates[rand.Intn(len(dates))],
			Title:  "Контент, который нужно опубликовать",
			User:   users[rand.Intn(len(users))],
			Places: []string{"NA", "MC"},
		}

		cards = append(cards, card)
	}

	err := layout.Execute(w, cards)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
