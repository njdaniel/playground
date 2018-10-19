package main

import "fmt"

func main() {
	input := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(input))
}

func longestCommonPrefix(strs []string) string {
	prefix := make([]rune,0)

	var min = func (x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	v1 := []rune(strs[0])
	v2 := []rune(strs[1])
	for i := 0; i < len(strs); i++ {
		for j := 0; j < min(len(v1),len(v2)); j++ {
			if v1[j] == v2[j] {
				prefix = append(prefix, v1[j])
			}
		}
		v1 = prefix
		prefix = nil
		if (i+2) < len(strs) {
			v2 = []rune(strs[i+2])
		}

	}

	return string(v1)
}
