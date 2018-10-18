package main

import "fmt"

func main() {
	input := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(input))
}

func longestCommonPrefix(strs []string) string {
	prefix := make([]rune,0)
	for i, j := range strs {

		if j[i] != j[i]-1 {

		}
	}
	return string(prefix)
}
