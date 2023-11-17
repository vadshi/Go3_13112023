package db

import (
	"context"
	"log"
	"math/rand"
	"time"
)

var Times = []string{
	"07:00",
	"08:00",
	"09:00",
	"10:00",
	"11:00",
	"12:00",
	"13:00",
	"14:00",
	"15:00",
	"16:00",
	"17:00",
	"18:00",
	"19:00",
	"20:00",
	"21:00",
	"22:00",
	"23:00",
}

func Seed(store *Store) {

	type Card struct {
		Time   string
		Date   string
		Title  string
		User   string
		Client string
		Sites  []string
	}

	var clients = []string{
		"ebayhamster",
		"bevywidely",
		"foresttopmast",
		"decoratesynonymous",
		"nonchalantopinion",
		"inquiryimminent",
		"dreambananas",
		"curryalembic",
		"lamentableconfirm",
		"acaciataste",
		"perfumedwrack",
		"serieswallabyies",
		"arcanasundress",
		"sownderwealthy",
		"scallopsmourner",
		"heartymorning",
		"malletneck",
	}

	var users = []string{
		"women/3.jpg",
		"women/4.jpg",
		"men/1.jpg",
	}

	var sites = []string{
		"SM",
		"MC",
		"NA",
		"NC",
		"SP",
		"VK",
		"TGG",
		"TGP",
	}

	days := 60
	start := time.Now().AddDate(0, 0, -days/2)
	end := start.AddDate(0, 0, days)

	var dates = []string{}
	for d := start; d.After(end) == false; d = d.AddDate(0, 0, 1) {
		dates = append(dates, d.Format("2006-01-02"))
	}

	cards := []Card{}
	for _, d := range dates {
		for _, t := range Times {
			if rand.Intn(2) == 0 { // 0 - empy card
				// just skip it
			} else { // 1 - card with content
				var card Card
				sitesRand := []string{}
				for _, site := range sites {
					if rand.Intn(2) == 1 {
						sitesRand = append(sitesRand, site)
					}
				}
				if len(sitesRand) > 4 {
					sitesRand = sitesRand[:4]
				}

				card = Card{
					Time:  t,
					Date:  d,
					Title: "Реклама для @" + clients[rand.Intn(len(clients))],
					User:  users[rand.Intn(len(users))],
					Sites: sitesRand,
				}
				cards = append(cards, card)
			}
		}
	}

	for _, card := range cards {

		insertedCard, err := store.CreateCard(context.Background(), CreateCardParams{
			Time:   card.Time,
			Date:   card.Date,
			Title:  card.Title,
			Client: card.Client,
			User:   card.User,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(insertedCard)
	}

}
