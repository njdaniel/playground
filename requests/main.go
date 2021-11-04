package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	url := "https://esi.evetech.net"
	ancestriesURL := url + "/latest/universe/ancestries/?datasource=tranquility&language=en"
	resp, err := http.Get(ancestriesURL)
	if err != nil {
		fmt.Println("error getting url: ", err)
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("error reading response body", err)
	}
	fmt.Println(string(body))
}

func SecureToken() {
	AuthorizationURL := "https://login.eveonline.com/v2/oauth/token"
	resp, err := http.PostForm()
}
