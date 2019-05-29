package main

import "fmt"

func main() {
	var a, b int
	var c = &b
	switch *c {
	case a:
		fmt.Println("a")
	case b:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
}
