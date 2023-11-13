package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/*
В контексте приглашений на мероприятия, RSVP — это запрос подтверждения от приглашённого человека или людей.
RSVP — это акроним французской фразы Répondez s’il vous plaît,
означающей буквально «Будьте добры ответить» или «Пожалуйста, ответьте».
*/

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 5)

func loadTemplates() {
	// TODO - load templates here
	// There are 5 templates: welcome.html, form.html, list.html, thanks.html, sorry.html
	templateNames := [5]string{"welcome", "form", "list", "thanks", "sorry"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+"html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

// welcomeHandler handles root URL
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	templates["welcome"].Execute(w, nil)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	templates["list"].Execute(w, responses)
}

// type formData struct {
// 	*Rsvp
// 	Errors []string
// }

func formHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}


func main() {
	loadTemplates()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
