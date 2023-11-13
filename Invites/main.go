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
	templates["welcome"].Execute(w, nil)
}

// listHandler handles "/list" URL
func listHandler(w http.ResponseWriter, r *http.Request) {
	templates["list"].Execute(w, responses)
}

type formData struct {
	*Rsvp
	Errors []string
}

// formHandler handles "/form" URL
func formHandler(w http.ResponseWriter, request *http.Request) {
	// TODO
	// Результат работы - пустая форма
	// GET localhost/form 
	// Взять данные формы(из request), проверить что данные не пустые и добавить очередное приглашение в список 
	// POST localhost/form 
	if request.Method == http.MethodGet {
		templates["form"].Execute(w, formData{
			Rsvp: &Rsvp{}, Errors: []string{},
		})
	} else if request.Method == http.MethodPost {
		request.ParseForm() // Парсим данные из request'а и записываем в request.Form
		responceData := Rsvp {
			Name: request.FormValue("name"),
			Email: request.Form["email"][0],  // the same as above
			Phone: request.FormValue("phone"),
			WillAttend: request.FormValue("willattend") == "true",
		}

		errors := []string{}
		// Проверка значений полей формы. Пустые поля недопустимы.
		if responceData.Name == "" {
			errors = append(errors, "Please enter your name")
		}
		if responceData.Email == "" {
			errors = append(errors, "Please enter your email")
		}
		if responceData.Phone == "" {
			errors = append(errors, "Please enter your phone")
		}
		if len(errors) > 0{
			templates["form"].Execute(w, formData{
				Rsvp: &responceData, Errors: errors,
			})
		} else {
			responses = append(responses, &responceData)
			if responceData.WillAttend {
				templates["thanks"].Execute(w, responceData.Name)
			} else {
				templates["sorry"].Execute(w, responceData.Name)
			}
		}
	}

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
