package main

import (
	"github.com/itwars/scraping/meteo/providers/allosurf"
	"github.com/itwars/scraping/meteo/providers/meteofrance"
	"github.com/itwars/scraping/meteo/providers/surfforecast"
)

func main() {
	meteofrance.Ephemeride()
	meteofrance.Previsions()
	allosurf.InfoMarine()
	surfforecast.Forecast()
}
