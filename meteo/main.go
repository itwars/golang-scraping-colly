package main

import (
	"github.com/itwars/golang-scraping-colly/meteo/providers/allosurf"
	"github.com/itwars/golang-scraping-colly/meteo/providers/grandesmarees"
	"github.com/itwars/golang-scraping-colly/meteo/providers/meteofrance"
	"github.com/itwars/golang-scraping-colly/meteo/providers/surfforecast"
)

func main() {
	meteofrance.Ephemeride()
	meteofrance.Previsions()
	allosurf.InfoMarine()
	surfforecast.Forecast()
	grandesmarees.GrandesMarees()
}
