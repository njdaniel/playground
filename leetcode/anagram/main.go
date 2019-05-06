package main

import (
	"fmt"
)

func main() {
	A := []int{12, 12, 46, 32, 50}
	B := []int{50, 12, 32, 46, 12}
	fmt.Println(anagramMappings(A, B))
}

func anagramMappings(A []int, B []int) []int {
	//create two hash maps, and go through both a and b, should get a O(N) run time
	//int num, position
	ha := make(map[int][]int, 0)
	hb := make(map[int][]int, 0)
	np := [100]int{}

	for i := 0; i < len(A); i++ {
		if _, ok := ha[A[i]]; !ok {
			ha[A[i]] = []int{i}
		} else {
			ha[A[i]] = append(ha[A[i]], i)
		}

		if _, ok := hb[B[i]]; !ok {
			hb[B[i]] = []int{i}
		} else {
			hb[B[i]] = append(hb[B[i]], i)
		}

		if _, ok := ha[B[i]]; ok {
			//insert b position on np[a-postion]
			np[ha[B[i]][0]] = i
		}
		if _, ok := hb[A[i]]; ok {
			np[ha[A[i]][0]] = hb[A[i]][0]
		}
	}
	snp := np[:len(A)]
	return snp
}
