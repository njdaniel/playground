package main

import "fmt"

func main() {
	input := []string{"aac","cab","abb"}
	fmt.Println(longestCommonPrefix(input))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs[0]) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
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
			} else {
				break
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
