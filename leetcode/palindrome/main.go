package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 121
	fmt.Println(IsPalindrome(x))
}

func IsPalindrome(x int) bool {
	s := strconv.Itoa(x)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reverse := string(runes)
	if s == reverse {
		return true
	}
	return false
}
