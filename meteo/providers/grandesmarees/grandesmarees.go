// Package grandesmarees providing scrapping information from maree.info
package grandesmarees

import (
	"fmt"
	"log"
	"regexp"

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
		// Expression régulière pour cleaner tout ce bordel
		// fmt.Println(infosGM.Informations)
		re := regexp.MustCompile(`([[:alpha:]]+) ([[:digit:]]+) ([a-zA-Zéèàê]+) ([[:digit:]]+) - coefficients* ([[:digit:]]+)\s*\/*\s*([[:digit:]]*) voir les horaires ...`)
		jour := re.ReplaceAllString(infosGM.Informations[0], `$1`)
		nJour := re.ReplaceAllString(infosGM.Informations[0], `$2`)
		mois := re.ReplaceAllString(infosGM.Informations[0], `$3`)
		fmt.Printf("%s %s %s\n", jour, nJour, mois)
	})
	c.Visit("http://maree.info/17/coefficients?c=gm")
}
