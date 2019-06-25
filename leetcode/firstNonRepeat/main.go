package main

import (
	"fmt"
)

func main() {
	input := "abzabc"
	fmt.Println(firstUniqChar(input))
}


func firstUniqChar(s string) int {
	m := make(map[rune][]int, 0)
	u := -1
	//ss := make([]int, 0)
	// loop string and add to m and ss
	//when adding check if exist, remove non-unique from ss
	for i, r := range s {
		if _, ok := m[r]; !ok {
			m[r] = []int{i,}
			//ss = append(ss, i)
		} else {
			m[r] = append(m[r], i)
		}
	}
	for _, v := range m {
		if len(v) == 1 && ((v[0]<u && u != -1)|| u ==-1 ){
			u = v[0]
		}
	}

	return u
}