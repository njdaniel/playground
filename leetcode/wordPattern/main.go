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
	m := make(map[rune]string)
	pr := []rune(pattern)
	if len(pr) != len(ss) {
		return false
	}
	for i := 0; i < len(pr); i++ {
		if _, ok := m[pr[i]]; !ok {
			m[pr[i]] = ss[i]
		} else if ok && m[pr[i]] != ss[i]{
			return false
		}
	}

	match = true
	return match
}