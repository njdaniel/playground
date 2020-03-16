package main

import "fmt"

func main() {
	n := []int32{3, 2, 1}
	countSwap(n)
}

func countSwap(a []int32)  {
	count := 0
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				count++
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps. \n" +
		"First Element: %d \n" +
		"Last Element: %d \n", count, a[0], a[len(a)-1])
}