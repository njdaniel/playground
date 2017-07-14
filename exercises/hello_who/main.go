package main

import "fmt"

func main() {
	var name string
	fmt.Print("Please enter a name: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}
