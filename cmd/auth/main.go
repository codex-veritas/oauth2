package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("-- Auth server --")

	http.HandleFunc("/", home)
	panic(http.ListenAndServe("localhost:8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("auth server"))
}
