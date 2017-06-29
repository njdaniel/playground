package main

import "fmt"

// Straight forward breaks when cond is met
func main() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
}
