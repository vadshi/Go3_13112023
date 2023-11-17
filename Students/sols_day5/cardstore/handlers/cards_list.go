package handlers

import (
	"local/hw/db"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func (server *Server) cardsList(ctx *gin.Context) {

	pageQuery := ctx.DefaultQuery("page", "0")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		log.Fatal(err)
	}

	templateName := "index.html.tmpl"
	if page != 0 {
		templateName = "cards_list"
	}

	// 7 days data by default
	days := 6 // 6 + current day
	start := time.Now().AddDate(0, 0, page*7)
	end := start.AddDate(0, 0, days)

	var dates = []string{}

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d.Format("2006-01-02"))
	}

	args := db.ListCardsParams{
		Date1: dates[0],
		Date2: dates[1],
		Date3: dates[2],
		Date4: dates[3],
		Date5: dates[4],
		Date6: dates[5],
		Date7: dates[6],
	}

	cards, err := server.store.ListCards(ctx, args)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("ERROR"))
			return
		}
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("ERROR"))
		return
	}

	data := struct {
		Prev  int
		Page  int
		Next  int
		Cards []db.Card
		Dates []string
	}{}

	data.Dates = dates
	data.Page = page
	data.Next = page + 1
	data.Prev = page - 1

	hasCard := func(cards []db.Card, d string, t string) (bool, db.Card) {
		for _, card := range cards {
			if card.Date == d && card.Time == t {
				return true, card
			}
		}
		return false, db.Card{}
	}

	// Add empty cards
	for _, d := range dates {
		for _, t := range db.Times {

			contains, card := hasCard(cards, d, t)

			if contains { // card with content
				data.Cards = append(data.Cards, card)
			} else { // empy card
				data.Cards = append(
					data.Cards, db.Card{
						Time: t,
						Date: d,
					})
			}
		}
	}

	ctx.HTML(http.StatusOK, templateName, data)
}
