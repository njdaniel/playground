package main

import "regexp"

func main() {

}

func isPalindrome(s string) bool {
	//convert to all lowercase
	//remove non-alphanumeric characters eg spaces, dashes
	//convert to []rune
	//loop through comparing first and last
	//ns := strings.ToLower(s)
	//ss = strings.Split(s, " ")

	//assuming all lowercase and alphanumeric
	processed := s
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	reverse := string(rs)
	if reverse == processed {
		return true
	}
	return false

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
