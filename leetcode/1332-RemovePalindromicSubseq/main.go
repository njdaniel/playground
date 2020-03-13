package main

import (
	"fmt"
	"playground/leetcode/palindrome/palindrome"
)

func main() {
	s := "bbaabaaa"
	fmt.Println(removePalindromeSub(s))
}

func removePalindromeSub(s string) int {
	pals := 0
	for x := len(s); x>0; x-- {
		 if palindrome.IsPalindrome([]rune(s[:x])) {
		 	pals++
		 	if x >= len(s) {
		 		s = ""
		 		break
			}
		 	s = s[x:]
		 	fmt.Println("new s: ", s)
		 	x = len(s)+1
		 }
	}
	return pals
}