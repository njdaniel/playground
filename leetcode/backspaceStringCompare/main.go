package main

import "fmt"

func main() {
	input1 := "abc"
	input2 := "abc"
	fmt.Println(backspaceCompare(input1, input2))
}

// backspaceCompare takes two strings and compares if they are equal given
// that '#' is a backspace and deletes the last char
func backspaceCompare(S, T string) bool {
	stackS := make([]rune, 0)
	stackT := make([]rune, 0)

	// push and pop onto stack for S
	for i, v := range S {
		if v == '#' {
			if len(stackS) > 0 {
				stackS = stackS[:i]
			}
		}
		stackS = append(stackS, v)
	}
	// push and pop onto stack for T
	for i, v := range T {
		if v == '#' {
			if len(stackT) > 0 {
				stackT = stackT[:i]
			}
		}
		stackT = append(stackT, v)
	}
	// compare if two strings are the same
	if string(stackS) == string(stackT) {
		return true
	}
	return false
}