package main

import (
	"local/hw1/db"
	"log"
	"math/rand"
)

var times = []string{
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

var dates = []string{
	"2023-11-13",
	"2023-11-14",
	"2023-11-15",
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

type Card struct {
	Time   string
	Date   string
	Title  string
	User   string
	Client string
	Sites  []string
}

func seed() {
	cards := []Card{}
	for _, d := range dates {
		for _, t := range times {

			var card Card
			if rand.Intn(2) == 0 { // 0 - empy card
				card = Card{
					Time: t,
					Date: d,
				}
			} else { // 1 - card with content

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
			}
			cards = append(cards, card)
		}
	}

	for _, card := range cards {
		insertedCard, err := queries.CreateCard(ctx, db.CreateCardParams{
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
