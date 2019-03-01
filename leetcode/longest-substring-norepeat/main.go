package main

import "fmt"

func main() {
	nput := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(nput))
}

func lengthOfLongestSubstring(s string) int {
	//create map storing runes to detect dupes,
	// if a dupe is found clear map and insert
	// the new rune
	var m = make(map[rune]rune, 0)

	//loop through s put rune in map,
	for _, r := range s {
		_, ok := m['r']
		if ok {
			for k := range m {
				delete(m, k)
			}
			m[r] = r
		} else {
			m[r] = r
		}
	}

	return len(m)
}
