package main

import "fmt"

func main() {
	input1 := "ab##c"
	input2 := "c"
	fmt.Println(backspaceCompare(input1, input2))
}

// backspaceCompare takes two strings and compares if they are equal given
// that '#' is a backspace and deletes the last char
func backspaceCompare(S, T string) bool {
	stackS := make([]rune, 0)
	stackT := make([]rune, 0)

	// push and pop onto stack for S
	for _, v := range S {
		if v == '#' {
			if len(stackS) > 0 {
				stackS = stackS[:(len(stackS)-1)]
			}
		} else {
			stackS = append(stackS, v)
		}
	}
	// push and pop onto stack for T
	for _, v := range T {
		if v == '#' {
			if len(stackT) > 0 {
				stackT = stackT[:(len(stackT)-1)]
			}
		} else {
			stackT = append(stackT, v)
		}
	}
	// compare if two strings are the same
	if string(stackS) == string(stackT) {
		return true
	}
	return false
}