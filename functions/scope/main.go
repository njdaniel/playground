package main

import "fmt"

func main() {
	count()
	count()
}

func count() int {
	x := 0
	fmt.Printf("init x = %d \n", x)
	f := func() {
		x++
		fmt.Printf("inside x = %d \n", x)
	}
	f()
	fmt.Printf("outside x = %d \n", x)
	return x
}
