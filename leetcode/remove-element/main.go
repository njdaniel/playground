package main

import "fmt"

func main() {
	n := []int{3, 2, 2, 3}
	fmt.Println(removeElement(n, 3))
	fmt.Println(n)
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}
