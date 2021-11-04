package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	getAncestries()
}

func secureToken(user, secret string) {
	authorizationURL := "https://login.eveonline.com/v2/oauth/authorize"
	v := url.Values{}
	v.Set("user", user)
	v.Set("password", secret)
	resp, err := http.PostForm(authorizationURL, v)
	if err != nil {
		fmt.Println("error authorizing: ", err)
	}
	if resp.StatusCode > 299 {
		fmt.Println("error status code: ", resp.Status)
	}
}

func getAncestries() {
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
