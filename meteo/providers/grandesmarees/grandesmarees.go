// Package grandesmarees providing scrapping information from maree.info
package grandesmarees

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type grandesmarees struct {
	Informations []string `selector:"li"`
}

// GrandesMarees func récupère les informations qui sont sous la forme d'un json
func GrandesMarees() {
	c := colly.NewCollector(
		colly.AllowedDomains("maree.info"),
	)
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Houston nous avons un problème avec Maree Info: ", err)
	})
	c.OnHTML(".ListeResultats", func(e *colly.HTMLElement) {
		infosGM := &grandesmarees{}
		e.Unmarshal(infosGM)
		fmt.Println("Date                  : ", infosGM.Informations)
	})
	c.Visit("http://maree.info/17/coefficients?c=gm")
}
