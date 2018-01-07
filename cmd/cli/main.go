package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// 1. Ask for User Authorization Grants

	var userName, password string
	fmt.Print("UserName: ")
	fmt.Scanln(&userName)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	// 2. Contact AuthServer to get a Token

	authResp, err := http.Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}
	defer authResp.Body.Close()
	authBody, _ := ioutil.ReadAll(authResp.Body)

	token := string(authBody)

	// 3. Send token to API Server to get data

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8081/dogs", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Token", token)

	apiResp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer apiResp.Body.Close()

	apiBody, _ := ioutil.ReadAll(apiResp.Body)

	fmt.Println(string(apiBody))
}
