package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern := "abab"
	str := "cat dog cat dog"
	fmt.Printf("Returned: %t", wordPattern(pattern, str))
	
}

func wordPattern(pattern string, str string) bool {
	match := false
	//ss := make([]string, 0)
	ss := strings.Split(str, " ")
	fmt.Printf("len ss: %d\n", len(ss))
	m := make(map[string]rune)
	m2 := make(map[rune]string)
	pr := []rune(pattern)
	if len(pr) != len(ss) {
		return false
	}
	for i := 0; i < len(pr); i++ {
		_, ok := m[ss[i]]
		_, ok2 := m2[pr[i]]
		if !ok && !ok2 {
			m[ss[i]] = pr[i]
			m2[pr[i]] = ss[i]
		} else if ok && ok2 && m[ss[i]] == pr[i]{
			continue
		} else {
			return false
		}
	}

	match = true
	return match
}