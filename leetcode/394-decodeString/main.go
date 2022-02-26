package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//s := "3[a]2[bc]"
	//s := "2[abb]"
	s := "2[abc]3[cd]ef"
	fmt.Println("Final: " + decodeString(s))
}

//topics: string, stack, recursion
//assumptions/contraints:
//  1 <= s.length <= 30
//  s consists of lowercase English letters, digits, and square brackets '[]'.
//  s is guaranteed to be a valid input.
//  All the integers in s are in the range [1, 300]
func decodeString(s string) string {
	//mutable temp of s, decoding replaces
	rs := []rune(s)

	//find simple section to decode
	//  if one bracket: decode
	//  no brackets: return
	//  multiple bracket pairs: simplify

	numbrackets := 0
	for _, v := range rs {
		if v == ']' {
			numbrackets++
		}
	}

	switch {
	case numbrackets > 1:
		fmt.Println("multiple brackets")
		//find the ] bracket and cut and replace the simplified encode string
		simp := make([]rune, 0)
		pbegin := 0
		pend := 0
		for i, v := range rs {
			if v == ']' {
				reachedBegin := true
				pend = i
				for j := i - 1; j >= 0; j-- {
					if rs[j] == ']' {
						simp = rs[j+1 : i+1]
						reachedBegin = false
						pbegin = j + 1
						break
					}
				}
				if reachedBegin {
					simp = rs[:i+1]
				}
				break
			}
		}
		//recursive get decodestring and cut and paste replacing the section
		fmt.Println("Simple: " + string(simp))
		simp = []rune(decodeString(string(simp)))
		newpend := pend - len(rs[pbegin:pend])
		rs = append(rs[:pbegin], rs[pend+1:]...)
		rs = append(rs[:pbegin], append(simp, rs[newpend:]...)...)
		fmt.Println("replaced: " + string(rs))
		//recursive
		rs = []rune(decodeString(string(rs)))
	case numbrackets == 1:
		fmt.Println("one bracket pair")
		k := 0
		krs := make([]rune, 0)
		encoded := make([]rune, 0)
		pbegin := 0
		pend := 0
		pEncodeBegin := 0
		for i, v := range rs {
			if v == '[' {
				for j := i - 1; j >= 0; j-- {
					if isLetter(rs[j]) {
						pbegin = j + 1
						break
					}
				}
				krs = rs[pbegin:i]
				kt, err := strconv.Atoi(string(krs))
				if err != nil {
					log.Fatal(err)
				}
				k = kt
				//encoded = rs[i+1 : len(rs)-1]
				pEncodeBegin = i + 1
				//break
			}
			if v == ']' {
				pend = i

				encoded = rs[pEncodeBegin:pend]
				break
			}
		}
		fmt.Printf("k: %d \t encoded: %v\n", k, encoded)
		//return string(decode(k, encoded))
		simp := string(rs[pbegin : pend+1])
		fmt.Println("SIMP: " + simp)
		dsimp := decode(k, encoded)
		newpend := pend - len(rs[pbegin:pend])
		rs = append(rs[:pbegin], rs[pend+1:]...)
		rs = append(rs[:pbegin], append(dsimp, rs[newpend:]...)...)
		//if pbegin != 0 {
		//	newpend := pend - len(rs[pbegin:pend])
		//	rs = append(rs[:pbegin], rs[pend+1:]...)
		//	rs = append(rs[:pbegin], append(dsimp, rs[newpend:]...)...)
		//} else {
		//	rs = decode(k, encoded)
		//}

	default:
		fmt.Println("no brackets")
	}

	fmt.Println("result: " + string(rs))
	return string(rs)
}

func decode(k int, encoded []rune) []rune {
	rs := make([]rune, 0)
	for i := 0; i < k; i++ {
		rs = append(rs, encoded...)
	}
	return rs
}

//func decodeString(s string) string {
//
//	//convert to []
//	//loop through s finding the end of the decodeString
//	//if the []rune isnt empty call decodeString
//	//else if its empty find the k and the s and create the []rune
//	//check if the []rune(s) is empty
//	//if yes return the result if not find the next part of the pattern
//	ra := []rune(s)
//	//raTemp := make([]rune, 0)
//
//	//check if only one bracket set
//	numbrackets := 0
//	for _, v := range ra {
//		if v == ']' {
//			numbrackets++
//		}
//	}
//	if numbrackets == 1 {
//		//k := 0
//		ks := make([]rune, 0)
//		es := make([]rune, 0)
//		p := 0
//		for i, v := range ra {
//			if v == '[' {
//				p = i
//				break
//			}
//			ks = append(ks, v)
//		}
//		k, err := strconv.Atoi(string(ks))
//		if err != nil {
//			log.Fatal(err)
//		}
//		es = ra[p : len(ra)-1]
//		decoded := make([]rune, 0)
//		for i := 0; i < k; i++ {
//			decoded = append(decoded, es...)
//		}
//	} else {
//
//		var pend int
//		var pbegin int
//		for i, v := range ra {
//			if v == ']' {
//				pend = i
//				//find the beginning of the k[encode](pbegin)
//				for j := i - 1; j >= 0; j-- {
//					if isNumber(ra[j]) {
//						for k := j - 1; k >= 0; k-- {
//							if !isNumber(ra[k]) {
//								pbegin = k + 1
//								break
//							}
//							pbegin = 0
//						}
//					}
//				}
//				pend = i + 1
//				raTemp := ra[pbegin:pend]
//				decodedString := decodeString(string(raTemp))
//				//cut + replace
//				ra = append(ra[:pbegin], ra[pend-1:]...)
//				ra = append(ra[:pbegin], append([]rune(decodedString), ra[pbegin:]...)...)
//				//break
//				i = 0
//			}
//		}
//	}
//
//	//remove and replace
//	return string(ra)
//}

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

func isLetter(r rune) bool {
	if r < 'a' || r > 'z' {
		return false
	}
	return true
}

func isNumber(r rune) bool {
	if r < '0' || r > '9' {
		return false
	}

	return true
}
