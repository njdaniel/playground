package main

import "fmt"

func main() {
	age := 44
	fmt.Println(&age)
	changeMe(&age)
	fmt.Println(age)
}

func changeMe(z *int) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = 24
	fmt.Println(z)
	fmt.Println(*z)
}

// Passed by reference?
// NOPE. Everything in Golang is passed by value
// BUT the value being passed could be a reference type
// NOT A REFERENCE TYPE is a primitive number
