package main

import (
	"regexp"
	"strings"
)

func main() {

}

//BigO: time:O(n) space:O(n)
func isPalindrome(s string) bool {
	//convert to all lowercase
	//remove non-alphanumeric characters eg spaces, dashes
	//convert to []rune
	//loop through comparing first and last
	//ns := strings.ToLower(s)
	//ss = strings.Split(s, " ")

	//assuming all lowercase and alphanumeric
	alphanumericString, _ := removeNonAlphaNumeric(s)
	lowercase := strings.ToLower(alphanumericString)
	rs := []rune(lowercase)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	reverse := string(rs)
	if reverse == lowercase {
		return true
	}
	return false

}

func isPalindromeOptimized(s string) bool {

	alphanumericString, _ := removeNonAlphaNumeric(s)
	lowercase := strings.ToLower(alphanumericString)
	rs := []rune(lowercase)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		if rs[i] != rs[j] {
			return false
		}
	}
	return true
}

//removeNonAlphaNumeric takes a string and removes all non-alphanumeric characters then returns new string
func removeNonAlphaNumeric(s string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	alphanumericString := reg.ReplaceAllString(s, "")
	return alphanumericString, nil

}
