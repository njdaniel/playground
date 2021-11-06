package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file := "phonebook.txt"
	phonebook := make([]string, 21)
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("error opening %s %v", phonebook, err)
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		phonebook = append(phonebook, scanner.Text())
	}
	search := ""
	results := make([]string, 0)
	fmt.Printf("Enter name or partal to search for: ")
	fmt.Scanln(&search)
	for _, v := range phonebook {
		if strings.Contains(v, search) {
			results = append(results, v)
		}
	}
	fmt.Println(results)
}
