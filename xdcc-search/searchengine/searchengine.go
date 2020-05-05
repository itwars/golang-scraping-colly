package searchengine

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// XdcceuResults exportation pour voir
type XdcceuResults struct {
	Network string
	Channel string
	Bot     string
	Slot    int
	Get     string
	Size    string
	Name    string
}

// Xdcceu function scrape le site xdcc.eu
func Xdcceu(item string) []XdcceuResults {
	itemEsc := strings.Replace(item, " ", "+", -1)
	c := colly.NewCollector(
		colly.AllowedDomains("www.xdcc.eu", "xdcc.eu"),
	)
	results := make([]XdcceuResults, 0, 10)
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Je visite le site suivant : ", r.URL.String())
	})
	c.OnHTML("body", func(e *colly.HTMLElement) {
		var index int = 0
		e.ForEachWithBreak("table tbody tr", func(_ int, el *colly.HTMLElement) bool {
			var net = strings.ToLower(el.ChildText("td:nth-of-type(1)"))
			var cha = strings.ToLower(el.ChildText("td:nth-of-type(2)"))
			var slo, _ = strconv.Atoi(strings.TrimPrefix(el.ChildText("td:nth-of-type(4)"), "#"))
			result := XdcceuResults{
				Network: net,
				Channel: cha,
				Bot:     el.ChildText("td:nth-of-type(3)"),
				Slot:    slo,
				Get:     el.ChildText("td:nth-of-type(5)"),
				Size:    el.ChildText("td:nth-of-type(6)"),
				Name:    el.ChildText("td:nth-of-type(7)"),
			}
			results = append(results, result)
			index++
			if index >= 1 {
				return false
			}
			return true
		})
	})
	c.OnScraped(func(s *colly.Response) {
	})
	//	c.Visit("https://www.xdcc.eu/search.php?searchkey=Altered+Carbon+S01+1080")
	c.Visit("https://www.xdcc.eu/search.php?searchkey=" + itemEsc)
	return results
}
