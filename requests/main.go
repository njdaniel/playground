package main

import (
	"fmt"
	"net/http"
)

func main() {

	url := "https://esi.evetech.net"
	ancestriesURL := url + "latest/universe/ancestries/?datasource=tranquility&language=en"
	resp, _ := http.Get(ancestriesURL)
	fmt.Println(*resp)
}
