package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codex-veritas/oauth2/pkg/jwt"
)

var secret = "secret"

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
	token := r.Header.Get("Token")
	payload, err := jwt.Decode(token, secret)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("cannot decode JWT"))
		return
	}

	user := payload.UserName
	fmt.Println("User:", user)
	// Retrieve dogs only for the given user

	// TODO filter dogs
	dogs := []dog{
		dog{"Snoopy", "Border colley", "blue"},
		dog{"James", "Terrier", "black"},
		dog{"Lily", "Terrier", "red"},
		dog{"Terrence", "Bichon Maltais", "pink"},
	}

	e := json.NewEncoder(w)
	err = e.Encode(dogs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Woof woof..."))
	}
}
