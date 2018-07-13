package main

import (
	"fmt"
)

func main() {
	s := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	fmt.Println(s)
	lenS := len(s)
	lenNew := lenS/2 + 1
	fmt.Println(lenS)
	fmt.Println(lenNew)
	newS := make([][]string, lenNew)

	// for i := 1; i < 8; i++ {
	// 	fmt.Println(s[:2])
	// 	if cap(s) >= 4 {
	// 		s = s[2:4]
	// 	} else {
	// 		s = s[2:3]
	// 	}

	// }

	for i := range newS {
		start := i * 2
		end := start + 2
		if end > lenS {
			end = lenS
		}
		newS[i] = s[start:end]
	}
	for i, v := range newS {
		fmt.Println(i, ":", v)
	}
}
