package main

import "fmt"

func main() {
	A := []int{12, 12, 46, 32, 50}
	B := []int{50, 12, 32, 46, 12}
	fmt.Println(anagramMappings(A, B))
}

func anagramMappings(A []int, B []int) []int {
	np := [100]int{}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				np[i] = j
			}
		}
	}
	snp := np[:len(A)]
	return snp
}
