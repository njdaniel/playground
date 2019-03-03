package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
}

//shouldnt this be passed as a pointer?
//also nums1 should be an array not slice, ifs its
//a set array
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1Copy := nums1
	fmt.Println(m, n, nums2, nums1Copy)
}
