package main

import "fmt"

func main() {
	n := []int{1}
	fmt.Println(removeElement(n, 1))
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	if val > len(nums)-1 {
		return len(nums)
	}
	nums = append(nums[:val], nums[val+1:]...)
	return len(nums)
}
