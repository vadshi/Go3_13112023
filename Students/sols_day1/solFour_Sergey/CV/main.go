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
	Name, Email, Phone, Post, Money string
}

type Works struct {
	Name, Post, Years string
}

// var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templateNames := [3]string{"welcome", "skills", "oldwork"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
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

	templates["welcome"].Execute(w, &Rsvp{
		Name: "Зверков Сергей Владимирович", 
		Email: "Zverkov@mail.ru", 
		Phone: "+7(999)999-99-99", 
		Post: "Император галактики", 
		Money: "1000000000000000000000000000000000000000$",
	})
}

// listHandler handles "/list" URL
func skillsHandler(w http.ResponseWriter, r *http.Request) {
	skills := []string{}
	skills = append(skills, "Управление солнечной системой")
	skills = append(skills, "Праздная жизнь")
	skills = append(skills, "Траты 1000000000$ в секунду")
	templates["skills"].Execute(w, skills)
}

func oldworkHandler(w http.ResponseWriter, r *http.Request) {
	work := []Works{}
	work = append(work, Works{Name: "Император солнечной системы", Post: "Император", Years: "500"})
	work = append(work, Works{Name: "Император Земли", Post: "Император", Years: "200"})
	work = append(work, Works{Name: "Российский царь", Post: "Цары", Years: "100"})
	templates["oldwork"].Execute(w, work)
}

func main() {
	loadTemplates()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/skills", skillsHandler)
	http.HandleFunc("/oldwork", oldworkHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
