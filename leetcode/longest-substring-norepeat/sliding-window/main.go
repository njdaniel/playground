package main

import (
	"fmt"
)

func main() {
	nput := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(nput))
}

func lengthOfLongestSubstring(s string) int {
	//first attempt was with brute force
	//it was slow and a little messy to
	longest := 0
	//each char has a index
	var m = make(map[rune]int, 0)
	for i, j := 0, 0; j <= len(s)-1 && i <= len(s); {
		if _, ok := m[rune(s[j])]; !ok {
			m[rune(s[j])] = j
			if j-i+1 > longest {
				longest = j - i + 1
			}
			j++
		} else {
			delete(m, rune(s[i]))
			i++
		}
	}
	return longest
}
