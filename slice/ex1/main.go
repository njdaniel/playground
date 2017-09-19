package main

import "fmt"

func main() {
	slice := []int{}
	fmt.Printf("%T\n", slice)
	make_slice := make([]int, 6)
	fmt.Printf("%T\n", make_slice)
	fmt.Println(len(make_slice))
	var s []int
	fmt.Printf("%T\n", s)
	array := [10]int{}
	fmt.Printf("%T\n", array)
	fmt.Println(len(array))
}
