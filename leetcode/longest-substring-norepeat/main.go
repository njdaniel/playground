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
	for idx, r := range s {
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
			lastrIdx := 0
			for i := idx - 1; i >= 0; i-- {
				// fmt.Printf("%T", s[i])
				if rune(s[i]) == r {
					lastr = true
					lastrIdx = i
				}
				if lastr {
					key := rune(s[i])
					dupe := false
					for j := lastrIdx + 1; j <= len(s)-1; j++ {
						if rune(s[j]) == key {
							dupe = true
						}
					}
					if !dupe {
						delete(m, key)
					}
				}
			}
			m[r] = r
			// for _, sub := range m {
			// 	fmt.Printf("%s", string(m[sub]))
			// }
			if len(m) > longest {
				longest = len(m)
			}
		} else {
			m[r] = r
			// fmt.Println(string(r))
			if len(m) > longest {
				longest = len(m)
			}
		}
	}
	// for _, sub := range m {
	// 	fmt.Printf("%s", string(m[sub]))
	// }
	return longest
}
