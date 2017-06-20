package main

import "fmt"

// Showing memory address of variable
func main() {
	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("%d \n", &a)
}

