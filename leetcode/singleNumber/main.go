package main

import "fmt"

func main() {
	input := []int{2, 2, 1}
	fmt.Println(singleNumber(input))
}

// given an non-empty array, every num appears twice except one, return that one
func singleNumber(nums []int) int {
	// make buckets
	b1 := map[int]struct{}{}
	b2 := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := b1[v]; ok {
			//remove from 1
			delete(b1, v)
			// add to bucket2
			b2[v] = struct{}{}
		} else {
			b1[v] = struct{}{}
		}
	}
	single := 0
	for k, _ := range b1 {
		single = k
		break
	}
	return single
}
