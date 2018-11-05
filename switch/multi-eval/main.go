package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Daniel", "Jenny":
		fmt.Println("Hi Daniel or Jenny")
	case "Medhi":
		fmt.Println("Hi Medhi")

	case "bob":
		fmt.Println("Hi boby")
	default:
		fmt.Println("You have no friends")
	}
}

/*
no default fallthrough
no break needed, break is automatic
but you can add a fallthrough to automatically call the next case
or have multiple
Prints:
	Hi Jenny
*/
