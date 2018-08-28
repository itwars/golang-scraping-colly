// Package rugbyrama providing scrapping information from rugbyrama.fr
package rugbyrama

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type top14 struct {
	Classement string `selector:".standing-table__cell--position"`
	Club       string `selector:".standing-table__cell--team"`
	NbPoints   string `selector:".standing-table__cell-value--main"`
}

// Top14 func récupère les informations qui sont sous la forme d'un json
func Top14() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.rugbyrama.fr", "rugbyrama.fr"),
	)
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Houston nous avons un problème avec Rugbyrama: ", err)
	})
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Je visite le site suivant : ", r.URL.String())
	})
	c.OnHTML("tr.standing-table__row", func(e *colly.HTMLElement) {
		classement := &top14{}
		e.Unmarshal(classement)
		fmt.Printf("%2s %-25s %4s point(s)\n", classement.Classement, classement.Club, classement.NbPoints)
	})
	// En fin de scraping j'affiche le json avec toutes les informations récupérées
	c.OnScraped(func(s *colly.Response) {
	})
	c.Visit("https://www.rugbyrama.fr/rugby/top-14/standing.shtml")
}
