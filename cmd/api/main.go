package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("-- API server --")

	http.HandleFunc("/", home)
	http.HandleFunc("/dogs", getDogs)
	panic(http.ListenAndServe("localhost:8081", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API server"))
}

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Color string `json:"color"`
}

func getDogs(w http.ResponseWriter, r *http.Request) {
	dogs := []dog{
		dog{"Snoopy", "Border colley", "blue"},
		dog{"James", "Terrier", "black"},
		dog{"Lily", "Terrier", "red"},
		dog{"Terrence", "Bichon Maltais", "pink"},
	}

	e := json.NewEncoder(w)
	err := e.Encode(dogs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Woof woof..."))
	}
}
