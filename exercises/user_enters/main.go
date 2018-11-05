package main

import "fmt"

func main() {
	var numOne int
	var numTwo int
	fmt.Print("Please enter a large number: ")
	fmt.Scan(&numOne)
	fmt.Print("Please enter a small number: ")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "/", numTwo, "=", float32(numOne)/float32(numTwo))
}
