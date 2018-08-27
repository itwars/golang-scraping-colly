// Package allosurf providing scrapping information from allosurf
package allosurf

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type marine struct {
	Jour []struct {
		Date  string `json:"date"`
		Daily struct {
			PortID       string `json:"port_id"`
			LowAm        string `json:"low_am"`
			HighAm       string `json:"high_am"`
			CoeffAm      string `json:"coeff_am"`
			LowPm        string `json:"low_pm"`
			HighPm       string `json:"high_pm"`
			CoeffPm      string `json:"coeff_pm"`
			Firstlight   string `json:"firstlight"`
			Lastlight    string `json:"lastlight"`
			Sunrise      string `json:"sunrise"`
			Sunset       string `json:"sunset"`
			Sst          string `json:"sst"`
			UvsIndex     string `json:"uvs_index"`
			UvsDesc      string `json:"uvs_desc"`
			UvsGlyphicon string `json:"uvs_glyphicon"`
			UvsColor     string `json:"uvs_color"`
		}
	}
}

// InfoMarine func récupère les informations qui sont sous la forme d'un json
func InfoMarine() {
	c := colly.NewCollector(
		colly.AllowedDomains("allosurf.net", "www.allosurf.net"),
	)
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Houston nous avons un problème avec Allosurf: ", err)
	})
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Je visite le site suivant : ", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		infosMarines := &marine{}
		json.Unmarshal(r.Body, infosMarines)
		fmt.Println("Biscarrosse")
		fmt.Println("Index UV              : ", infosMarines.Jour[0].Daily.UvsIndex)
		fmt.Println("Coefficient matin     : ", infosMarines.Jour[0].Daily.CoeffAm)
		fmt.Println("Coefficient après-midi: ", infosMarines.Jour[0].Daily.CoeffPm)
		fmt.Println("Pleine mer matin      : ", infosMarines.Jour[0].Daily.HighAm)
		fmt.Println("Base mer matin        : ", infosMarines.Jour[0].Daily.LowAm)
		fmt.Println("Pleine mer après-midi : ", infosMarines.Jour[0].Daily.HighPm)
		fmt.Println("Basse mer après-midi  : ", infosMarines.Jour[0].Daily.LowPm)
	})
	// En fin de scraping j'affiche le json avec toutes les informations récupérées
	c.OnScraped(func(s *colly.Response) {
	})
	c.Visit("https://www.allosurf.net/meteo/spot/data.php?spoId=116&meteo=wam_1_wrf_5_96_h&start=1&end=1")
}
