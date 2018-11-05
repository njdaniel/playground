package main

import (
	"fmt"
)

func SumEvenFibonacci(limit int) int {
	fib1 := 1
	fib2 := 2
	var evenFib []int
	sum := 0
	for fib2 <= limit {
		var newFib int = fib1 + fib2
		fib1 = fib2
		fib2 = newFib
		evenFib = append(evenFib, newFib)
	}
	for i := 0; i >= len(evenFib); i++ {
		sum = evenFib[i]
	}
	return sum
}

func main() {
	bob := SumEvenFibonacci(8)
	fmt.Println(bob)
}
