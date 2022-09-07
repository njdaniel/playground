package main

import "fmt"

func main() {
	date := []int{0, 1}
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		if date[i] {
			fmt.Println("true")
		}
	}
}
