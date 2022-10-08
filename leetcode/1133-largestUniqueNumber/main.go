package main

import "fmt"

func main() {
	fmt.Print(largestUniqueNumber([]int{9, 9, 8}))
}

func largestUniqueNumber(nums []int) int {
	lUN := -1
	//brute force
	//eliminate dupes
	unique := make(map[int]int, 0)
	for _, v := range nums {
		if _, ok := unique[v]; !ok {
			unique[v] = 1
		} else {
			unique[v] = 0
		}
	}
	//find biggest numb
	for k, v := range unique {
		if k > lUN && v == 1 {
			lUN = k
		}
	}

	return lUN
}
