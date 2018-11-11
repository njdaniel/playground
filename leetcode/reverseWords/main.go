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
	for _, word := range words {
		rv := ""
		for _, l := range word {
			rv = string(l) + rv
		}
		if ra == ""{
			ra = rv
		} else {
			ra = ra + " " + rv
		}

	}
	return ra
}
