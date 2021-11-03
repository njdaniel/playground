package main

import (
	"fmt"
	"io"
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
	var body []byte
	_, err = io.ReadFull(resp.Body, body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("error io reading body: ", err)
	}
	fmt.Println(body)
}
