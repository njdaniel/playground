package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Let's take LeetCode contest"
	fmt.Println(reverseWords(input))
}

func reverseWords(s string) string {
	ra := ""
	//split by words
	words := strings.Fields(s)
	fmt.Println(words[0])
	for _, word := range words {
		rv := ""
		for _, l := range word {
			rv = string(l) + rv
		}
		ra = ra +" " + rv
	}
	return ra
}
