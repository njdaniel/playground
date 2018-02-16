package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Hi Daniel")
	case "Medhi":
		fmt.Println("Hi Medhi")
		fallthrough
	case "Jenny":
		fmt.Println("Hi Jenny")
	default:
		fmt.Println("You have no friendsd")
	}
}

/*
no default fallthrough
no break needed, break is automatic
but you can add a fallthrough to automatically call the next case
Prints:
	Hi Medhi
	Hi Jenny
*/
