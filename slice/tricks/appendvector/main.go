package main

import (
	"fmt"
)

func main() {
	b := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	var a []string
	a = append(a, b...)
	fmt.Println(a)
}
