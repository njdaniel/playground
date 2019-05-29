package main

import "fmt"

func wrapper() func() int {
	var x int
	//x := 0
	fmt.Printf("outside x=%d\n", x)
	return func() int {
		x++
		fmt.Printf("inside func x=%d\n", x)
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
