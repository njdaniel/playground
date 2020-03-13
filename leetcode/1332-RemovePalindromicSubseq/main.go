package main

import (
	"fmt"
	"playground/leetcode/palindrome/palindrome"
)

func main() {
	s := "abb"
	fmt.Println(removePalindromeSub(s))
}

func removePalindromeSub(s string) int {
	pals := 0
	for x := len(s); x>0; x-- {
		 if palindrome.IsPalindrome([]rune(s[:x])) {
		 	pals++
		 	s = s[x:]
		 	x = len(s)
		 }
	}
	return pals
}