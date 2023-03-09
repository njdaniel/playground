package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("in main")
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var m float64
	//create new slice
	combinedIntSlice := []int{}
	combinedIntSlice = append(combinedIntSlice, nums1...)
	combinedIntSlice = append(combinedIntSlice, nums2...)
	//sort
	sort.Ints(combinedIntSlice)

	//get the median
	// odd, even avg of 2 middle
	switch {
	case len(combinedIntSlice)%2 == 0:
		m = (float64(combinedIntSlice[len(combinedIntSlice)/2]) + float64(combinedIntSlice[len(combinedIntSlice)/2-1])) / 2.0
	case len(combinedIntSlice)%2 == 1:
		m = float64(combinedIntSlice[len(combinedIntSlice)/2])
	}
	return m
}
