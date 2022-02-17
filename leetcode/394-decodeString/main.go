package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	s := "3[a]2[bc]"
	fmt.Println(decodeString(s))
}

//topics: string, stack, recursion
//assumptions/contraints:
//  1 <= s.length <= 30
//  s consists of lowercase English letters, digits, and square brackets '[]'.
//  s is guaranteed to be a valid input.
//  All the integers in s are in the range [1, 300]
func decodeString(s string) string {

	//convert to []
	//loop through s finding the end of the decodeString
	//if the []rune isnt empty call decodeString
	//else if its empty find the k and the s and create the []rune
	//check if the []rune(s) is empty
	//if yes return the result if not find the next part of the pattern
	ra := []rune(s)
	//raTemp := make([]rune, 0)

	//check if only one bracket set
	numbrackets := 0
	for _, v := range ra {
		if v == ']' {
			numbrackets++
		}
	}
	if numbrackets == 1 {
		//k := 0
		ks := make([]rune, 0)
		es := ""
		p := 0
		for i, v := range ra {
			if v == '[' {
				p = i
				break
			}
			ks = append(ks, v)
		}
		k, err := strconv.Atoi(string(ks))
		if err != nil {
			log.Fatal(err)
		}
	} else {

		var pend int
		var pbegin int
		for i, v := range ra {
			if v == ']' {
				pend = i
				//find the beginning of the k[encode](pbegin)
				for j := i - 1; j >= 0; j-- {
					if isNumber(ra[j]) {
						for k := j - 1; k >= 0; k-- {
							if !isNumber(ra[k]) {
								pbegin = k + 1
								break
							}
							pbegin = 0
						}
					}
				}
				pend = i + 1
				raTemp := ra[pbegin:pend]
				decodedString := decodeString(string(raTemp))
				//cut + replace
				ra = append(ra[:pbegin], ra[pend-1:]...)
				ra = append(ra[:pbegin], append([]rune(decodedString), ra[pbegin:]...)...)
				//break
				i = 0
			}
		}
	}

	//remove and replace
	return string(ra)
}

//// Stack type for first in last out
//type Stack []rune
//
////IsEmpty method for checking if stack is empty
//func (s *Stack) IsEmpty() bool {
//	return len(*s) == 0
//}
//
////Pop method
//// return false if empty
//func (s *Stack) Pop() (rune, bool) {
//	if s.IsEmpty() {
//		return 0, false
//	}
//	e := s[len(s)-1]
//	s = s[:len(s)-1]
//	return e, true
//}
//
////checkDecodeInput verifies the input is simplified
////
//func checkDecodeInput(s string) bool {
//	rs := Stack(s)
//	var k int
//	var encodedString []rune
//	for i := 0; ; i++ {
//		if i == 0 {
//			continue
//		}
//
//	}
//}

func isLetter(s string) bool {
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}

func isNumber(r rune) bool {
	if r < '0' || r > '9' {
		return false
	}

	return true
}
