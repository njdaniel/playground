package palindrome

import "strconv"

// Goal:
// Use polymorphism / generics / substitute functions for different types for a function
// Given an array of different possible types (ex, int, rune)
// ..probably should just use rune. Convert any number, string, to []rune for this case.

type Palindrome interface {
	IsPalindrome() bool
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