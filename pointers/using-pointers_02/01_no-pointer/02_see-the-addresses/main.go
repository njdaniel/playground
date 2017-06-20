package main

import (
	"fmt"
)

// passing the VALUE of x
func zero(z int) {
	fmt.Printf("%p\n", &z) // address of func zero
	fmt.Println(&z) // address in func zero
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x) // address in main
	zero(x)
	fmt.Println(x) // x is still 5
}
