package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x2081a220

	var b *int = &a
	fmt.Println(b)  // 0x2081a220
	fmt.Println(*b) // 43

	// b is an int pointer
	// b points to teh memory address where int is stored
	// to see the value in the memory address, add a * in front
	// this is derefencing
}
