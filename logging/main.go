package main

import (
	"os"
	"log"
	"fmt"
)

func main() {
	f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	file, err := os.Open("test.txt")
	if err != nil {
		log.Panic("Panicing!")
	}
	defer file.Close()
	fmt.Println("Do I print?")
}
