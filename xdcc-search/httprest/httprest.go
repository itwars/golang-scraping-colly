package httprest

import (
	"bytes"
	"fmt"
	"net/http"
)

// RequestDL envoie une requête http REST vers une autre application qui "écoute" sur le port 8080
func RequestDL(item string) {
	var jsonStr = []byte(item)
	resp, err := http.Post("http://localhost:8080/", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		print(err)
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		resp.Body.Close()
	}
}
