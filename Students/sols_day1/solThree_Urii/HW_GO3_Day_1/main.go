package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type education struct {
	Name_education string
	Specialty      string
}

type experience struct {
	Company  string
	Position string
	Function string
	Begin    string
	End      string
}

type cvData struct {
	Name       string
	LastName   string
	Email      string
	Phone      string
	Position   string
	Education  []education
	Experience []experience
}

var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templateNames := [1]string{"cv_body"}
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

func showCV(w http.ResponseWriter, r *http.Request) {

	educations := []education{
		{
			Name_education: "ТюмГУ",
			Specialty:      "Экономика",
		},
		{
			Name_education: "МИФУБ",
			Specialty:      "Оценка бизнеса",
		},
	}

	user_experience := []experience{
		{
			Company:  "Банк Первомайский",
			Position: "Контролер",
			Function: "Контроль за соблюдением законодательства в сферах РЦБ, ПОД-ФТ, Инсайда",
			Begin:    "2012-01-01",
			End:      "2015-05-30",
		},
		{
			Company:  "УК Премьер Лига",
			Position: "Контролер УК",
			Function: "Контроль за соблюдением законодательства в сферах РЦБ, ПОД-ФТ, Инсайда",
			Begin:    "2015-06-01",
			End:      "2019-04-15",
		},
	}

	cv_data := cvData{
		Name:       "Jon",
		LastName:   "Libovsky",
		Email:      "mail@mail.ru",
		Phone:      "8-925-635-84-62",
		Position:   "Финансовый контролер",
		Education:  educations,
		Experience: user_experience,
	}

	templates["cv_body"].Execute(w, cv_data)
}

func main() {
	loadTemplates()
	http.HandleFunc("/", showCV)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
