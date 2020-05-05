package main

import (
	"encoding/json"
	"fmt"
	"os"
	"xdcc-search/httprest"
	"xdcc-search/searchengine"
)

func main() {
	searchItem := os.Args[1]
	results := make([]searchengine.XdcceuResults, 0, 10)
	results = searchengine.Xdcceu(searchItem)
	out, err := json.Marshal(results[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	httprest.RequestDL(string(out))
}
