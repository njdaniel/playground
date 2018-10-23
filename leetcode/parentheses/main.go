package main

import "fmt"

func main() {
	input := "()"
	fmt.Println(isValid(input))
}

func isValid(s string) bool {
	stack := make([]rune,0)
	rs := []rune(s)
	for _, j := range rs  {
		switch j {
		case '(', '{', '[':
			stack = append(stack, j)
		case ')':
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}