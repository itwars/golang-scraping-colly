// Package foot providing scrapping information from francefootball.fr
package foot

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type ligue1 struct {
	Rang     string `selector:".table__col--rank"`
	Club     string `selector:".table__col--team"`
	NbPoints string `selector:".table__col--pts"`
}

// Ligue1 func récupère les informations qui sont sous la forme d'un json
func Ligue1() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.francefootball.fr", "francefootball.fr"),
	)
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Houston nous avons un problème avec Francefootball: ", err)
	})
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Je visite le site suivant : ", r.URL.String())
		fmt.Println()
		fmt.Println("Classement foot ligue 1")
		fmt.Println()
	})
	c.OnHTML("tr.table__row", func(e *colly.HTMLElement) {
		classement := &ligue1{}
		e.Unmarshal(classement)
		fmt.Printf("%2s %-25s %4s point(s)\n", classement.Rang, classement.Club, classement.NbPoints)
	})
	// En fin de scraping j'affiche le json avec toutes les informations récupérées
	c.OnScraped(func(s *colly.Response) {
	})
	c.Visit("https://www.francefootball.fr/ligue-1/classement/")
}
