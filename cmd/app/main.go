package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	flags "github.com/jessevdk/go-flags"
)

type options struct {
	Bind string `long:"bind" default:"0.0.0.0" description:"address of the server"`
	Port int    `long:"port" default:"3000" description:"port of the server"`
}

func main() {
	opts := options{}
	flags.ParseArgs(&opts, os.Args)

	l := log.New(os.Stdout, "MAIN  ", log.LstdFlags|log.Lshortfile)

	http.HandleFunc("/", home(opts.Bind+":8080"))
	addr := opts.Bind + ":" + strconv.Itoa(opts.Port)
	l.Printf("listening on %v...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func home(authURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		signInLink := "http://" + authURL + "/authorize"
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
}
