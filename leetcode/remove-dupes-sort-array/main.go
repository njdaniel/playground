package main

import "fmt"

func main() {
	n := []int{1,1,2}
	fmt.Println(removeDuplicates(n))
}

func removeDuplicates(nums []int) int {
	limit := len(nums)-2
	for i := 0; i < limit; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}
