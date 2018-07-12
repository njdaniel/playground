package main

import (
	"fmt"
)

func main() {
	s := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	fmt.Println(s)

	for i := 1; i < 8; i++ {
		fmt.Println(s[:2])
		if cap(s) >= 4 {
			s = s[2:4]
		} else {
			s = s[2:3]
		}

	}
}
