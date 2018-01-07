package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("-- API server --")

	http.HandleFunc("/", home)
	panic(http.ListenAndServe("localhost:8081", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API server"))
}
