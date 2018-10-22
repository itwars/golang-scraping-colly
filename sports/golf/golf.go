// Package golf récupère les information sur le site du Figaro
package golf

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type leaderboard struct {
	Rang     string `selector:".div_idalgo_sport_standing_position"`
	Nom      string `selector:".div_idalgo_sport_standing_country > img" attr:"title"`
	Pays     string `selector:".div_idalgo_sport_standing_name"`
	NbPoints string `selector:".div_idalgo_sport_standing_total"`
}

// Leaderboard func récupère les informations de classement
func Leaderboard() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.lefigaro.fr", "golf.lefigaro.fr"),
	)
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Houston nous avons un problème avec Le Figaro: ", err)
	})
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Je visite le site suivant : ", r.URL.String())
		fmt.Println()
		fmt.Println("Classement mondiale de golf")
		fmt.Println()
	})
	c.OnHTML("[data-federation=\"9\"] ul.ul_idalgo_sport_standing > li", func(e *colly.HTMLElement) {
		classement := &leaderboard{}
		e.Unmarshal(classement)
		value, _ := strconv.ParseFloat(classement.NbPoints, 32)
		if value != 0 {
			fmt.Printf("%2s %-25s %-25s %3.2f point(s)\n", classement.Rang, classement.Nom, classement.Pays, value)
		}
	})
	// En fin de scraping j'affiche le json avec toutes les informations récupérées
	c.OnScraped(func(s *colly.Response) {
	})
	c.Visit("http://golf.lefigaro.fr/actualite/live/classement-mondial/1/classement")
}
