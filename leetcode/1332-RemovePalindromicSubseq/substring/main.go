package main

import (
	"fmt"
	"playground/leetcode/palindrome/palindrome"
)

func main() {
	s := "bbaabaaa"
	fmt.Println(removePalindromeSub(s))
}

// this is INCORRECT, this is for consecutive substring not subsequence
// this only removes palindromes that already exist in consecutive order and removes
// the substring, question was asking for subsequence which can remove characters inbetween
// but cannont be reordered
// also didnt realize that at most the response will be 2..
// ie remove all a's and remove all b's
// just added limit of 2 and that works
func removePalindromeSub(s string) int {
	pals := 0
	for x := len(s); x>0; x-- {
		 if palindrome.IsPalindrome([]rune(s[:x])) {
		 	pals++
		 	if pals > 2 {
		 		pals = 2
		 		return pals
			}
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