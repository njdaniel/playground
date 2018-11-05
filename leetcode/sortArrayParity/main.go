package main

import "fmt"

func main() {
	input := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(input))
}

func sortArrayByParity(A []int) []int {
	//return evens then odds
	even := make([]int, 0)
	odd := make([]int, 0)
	B := make([]int, 0)
	for _, v := range A {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	B = append(even, odd...)
	return B
}
