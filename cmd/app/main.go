package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	signInLink := "http://localhost:8080/authorize"
	title := "Welcome"
	body := fmt.Sprintf("Please sign in <br><a href=\"%s\">here</a>", signInLink)
	html := fmt.Sprintf("<html><head><title>%s</title></head><body>%s</body></html>", title, body)
	fmt.Fprint(w, html)

	// m := map[string]string{
	// 	"response_type": "",
	// 	"client_id":     "",
	// 	"redirect_uri":  "",
	// 	"scope":         "",
	// 	"state":         "",
	// }
}
