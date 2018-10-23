package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

type freelance struct {
	LesPostes []poste `json:"postes"`
}

type poste struct {
	Profil string `json:"profil"`
	TJM    int    `json:"tjm"`
}

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

func retResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func style(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "style.css")
}

func api(w http.ResponseWriter, r *http.Request) {
	info := map[string]int{}
	c := colly.NewCollector(
		colly.AllowedDomains("freelance-info.fr", "www.freelance-info.fr"),
	)
	c.OnError(func(_ *colly.Response, err error) {
		retResponse(w, http.StatusNotFound, map[string]string{"error": "colly in api func"})
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
	// En fin de scraping j'affiche le json avec toutes les informations récupérées
	c.OnScraped(func(s *colly.Response) {
		list := &freelance{}
		for profil, tjm := range info {
			p := poste{profil, tjm}
			list.LesPostes = append(list.LesPostes, p)
		}
		retResponse(w, http.StatusOK, list)
	})
	c.Visit("https://www.freelance-info.fr/tarifs/")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	retResponse(w, http.StatusNotFound, map[string]string{"error": "Not found"})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/api", api)
	router.HandleFunc("/style.css", style)
	router.NotFoundHandler = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000", router)
}
