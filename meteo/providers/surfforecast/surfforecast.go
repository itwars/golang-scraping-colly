// Package surfforecast providing scrapping information from surf-forecast
package surfforecast

import (
	"fmt"
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

type forecast struct {
	Jours        []string `selector:"#table-day > td strong"`
	HauteurHoule []string `selector:"#table-temp > td img" attr:"title"`
	Vent         []string `selector:"#table-wind > td img" attr:"alt"`
}

var jours = map[string]string{
	"Mon": "Lundi",
	"Tue": "Mardi",
	"Wed": "Mercredi",
	"Thu": "Jeudi",
	"Fri": "Vendredi",
	"Sat": "Samedi",
	"Sun": "Dimanche",
}

// Forecast func récupère les informations qui sont sous la forme d'un json
func Forecast() {
	c := colly.NewCollector(
		colly.AllowedDomains("fr.surf-forecast.com"),
	)
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Houston nous avons un problème avec Surf Forecast: ", err)
	})
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Je visite le site suivant : ", r.URL.String())
	})
	c.OnHTML("table", func(e *colly.HTMLElement) {
		infosForecast := &forecast{}
		e.Unmarshal(infosForecast)
		var re = regexp.MustCompile(`(\d+)\s+(\w+)`)
		var resultat = re.FindStringSubmatch(infosForecast.Vent[0])
		fmt.Println("Hauteur de houle      : ", infosForecast.HauteurHoule[0], "m")
		fmt.Println("Force du vent         : ", resultat[1])
		fmt.Println("Direction du vent     : ", resultat[2])
	})
	// En fin de scraping j'affiche le json avec toutes les informations récupérées
	c.OnScraped(func(s *colly.Response) {
	})
	c.Visit("http://fr.surf-forecast.com/breaks/Biscarosse-Plage/forecasts/feed/a")
}
