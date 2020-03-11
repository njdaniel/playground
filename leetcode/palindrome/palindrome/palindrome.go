package palindrome

import (
	"fmt"
	"reflect"
	"strconv"
)

// Goal:
// Use polymorphism / generics / substitute functions for different types for a function
// Given an array of different possible types (ex, int, rune)
// ..probably should just use rune. Convert any number, string, to []rune for this case.
// otherwise, the interface would need to use reflection to determine the type and changed to
// []rune anyways.
// still might be nice to use interface to abstract the types for determining if its a palindrome

type Palindrome interface {
	IsPalindrome() bool
}



func IsPalindromeOriginal(x int) bool {
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

func IsPalindrome(runes []rune) bool {
	reverse := make([]rune, len(runes))
	copy(reverse, runes)
	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	fmt.Println(reverse)
	// compare two arrays?
	if reflect.DeepEqual(runes, reverse) {
		fmt.Println("runes ", runes, " = ", reverse)
		return true
	}
	return false
}