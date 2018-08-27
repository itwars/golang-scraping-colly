package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
)

// PrintInfo réalise les tests et recherche de patterns afin de retourner l'intitulé du poste et le TJM
func PrintInfo(e *colly.HTMLElement) (string, int, bool) {
	var title string
	var salary int
	var err bool
	// Comme le site scrappé est mal fichu, j'élimine de ma recherche les <li> suivants
	matched, _ := regexp.MatchString("Freelance|Réflexion|Travail|Formation|Typologie|astreinte", e.Text)
	if !matched {
		// J'extrait l'intitulé et le montant du TJM
		var re = regexp.MustCompile(`\s*(Maj|Lu)?\s*(.*)\s+(\d*) €\/j`)
		title = re.ReplaceAllString(e.Text, `$2`)
		sal := re.ReplaceAllString(e.Text, `$3`)
		salary, _ = strconv.Atoi(sal)
		if salary == 0 {
			err = true
		}
	} else {
		err = true
	}
	return title, salary, err
}

func main() {
	info := map[string]int{}
	c := colly.NewCollector(
		colly.AllowedDomains("freelance-info.fr", "www.freelance-info.fr"),
	)
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Houston nous avons un problème : ", err)
	})
	c.OnHTML("li.v1", func(e *colly.HTMLElement) {
		t, s, err := PrintInfo(e)
		if !err {
			info[t] = s
		}
	})
	c.OnHTML("li.v2", func(e *colly.HTMLElement) {
		t, s, err := PrintInfo(e)
		if !err {
			info[t] = s
		}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Je visite le site suivant : ", r.URL.String())
	})
	// En fin de scraping j'affiche le json avec toutes les informations récupérées
	c.OnScraped(func(s *colly.Response) {
		jsonString, _ := json.Marshal(info)
		fmt.Printf("%s", jsonString)
	})
	c.Visit("https://www.freelance-info.fr/tarifs/")
}
