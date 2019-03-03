package main

import "fmt"

func main() {
	nput := "bpfbhmipx"
	fmt.Println(lengthOfLongestSubstring(nput))
}

func lengthOfLongestSubstring(s string) int {
	//create map storing runes to detect dupes,
	// if a dupe is found clear map and insert
	// the new rune
	var m = make(map[rune]rune, 0)
	longest := 0
	//loop through s put rune in map,
	for num, r := range s {
		_, ok := m[r]
		if ok {
			if len(m) > longest {
				longest = len(m)
			}
			// for k := range m {
			// delete(m, k)
			// }
			//remove (last ref to r)->0
			lastr := false
			for i := num - 1; i >= 0; i-- {
				// fmt.Printf("%T", s[i])
				if rune(s[i]) == r {
					lastr = true
				}
				if lastr {
					key := rune(s[i])
					delete(m, key)
				}
			}
			m[r] = r
		} else {
			m[r] = r
			if len(m) > longest {
				longest = len(m)
			}
		}
	}
	return longest
}
