package main

import (
	"fmt"
	"log"
	"os"

	"playground/gopl.io/ch5/links"
)

//alternate to crawl2 is to use just 20https requests concurrently
func main() {
	worklist := make(chan []string)  //list of urls, may have dupes
	unseenLinks := make(chan string) //de-dupe

	//add cl args to worklist
	go func() { worklist <- os.Args[1:] }()

	//create 20 crawlers goroutines to fetch each unseen link
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	//de-dupe worklist
	//send unseen to crawler
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // send blank to acquire token space
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
