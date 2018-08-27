// Package meteofrance providing scrapping information from meteofrance
package meteofrance

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type ephemeride struct {
	LeveSoleil   string `json:"heureLeveSoleil"`
	CoucheSoleil string `json:"heureCoucheSoleil"`
	LeveLune     string `json:"heureLeveLune"`
	CoucheLune   string `json:"heureCoucheLune"`
	Saint        string `json:"saint"`
	Date         string `json:"dateRef"`
}

// Ephemeride func récupère les informations qui sont sous la forme d'un json
func Ephemeride() {
	c := colly.NewCollector(
		colly.AllowedDomains("meteofrance.com", "www.meteofrance.com"),
	)
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Houston nous avons un problème avec Météo France: ", err)
	})
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Je visite le site suivant : ", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		infosEphemeride := &ephemeride{}
		json.Unmarshal(r.Body, infosEphemeride)
		fmt.Println("Date du jour          : ", infosEphemeride.Date)
		fmt.Println("Le soleil se lève à   : ", infosEphemeride.LeveSoleil)
		fmt.Println("Le soleil se couche à : ", infosEphemeride.CoucheSoleil)
		fmt.Println("Saint(e) du jour      : ", infosEphemeride.Saint)
	})
	// En fin de scraping j'affiche le json avec toutes les informations récupérées
	c.OnScraped(func(s *colly.Response) {
	})
	c.Visit("https://www.meteofrance.com/mf3-rpc-portlet/rest/ephemerides/DEPT33")
}

// Previsions récupère les informations de température à l'intérieur de différentes balises html
func Previsions() {
	c := colly.NewCollector(
		colly.AllowedDomains("meteofrance.com", "www.meteofrance.com"),
	)
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Je visite le site suivant : ", r.URL.String())
	})
	c.OnHTML("div.liste-jours li.active", func(e *colly.HTMLElement) {
		fmt.Println("Tendance du temps     : ", e.Attr("title"))
		fmt.Println("Température           : ", e.ChildText("span.min-temp"))
		fmt.Println("Température           : ", e.ChildText("span.max-temp"))
	})
	c.OnScraped(func(s *colly.Response) {
	})
	c.Visit("https://www.meteofrance.com/previsions-meteo-france/bordeaux/33000")
}
