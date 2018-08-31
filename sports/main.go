package main

import (
	"github.com/itwars/golang-scraping-colly/sports/foot"
	"github.com/itwars/golang-scraping-colly/sports/golf"
	"github.com/itwars/golang-scraping-colly/sports/rugby"
)

func main() {
	rugby.Top14()
	foot.Ligue1()
	golf.Leaderboard()
}
