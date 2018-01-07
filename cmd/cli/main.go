package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("cli")

	apiResp, err := http.Get("http://localhost:8081")
	if err != nil {
		panic(err)
	}
	defer apiResp.Body.Close()

	authResp, err := http.Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}
	defer authResp.Body.Close()

	apiBody, _ := ioutil.ReadAll(apiResp.Body)
	authBody, _ := ioutil.ReadAll(authResp.Body)

	fmt.Println(string(apiBody))
	fmt.Println(string(authBody))
}
