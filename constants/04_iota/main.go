package main

import "fmt"

const (
	A = iota // 0
	B = iota * 10// 1 * 10
	C = iota * 10 // 2 * 10
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

// iota def is a very small increment value