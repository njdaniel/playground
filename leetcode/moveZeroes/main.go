package main

import "fmt"

func main() {
	input := []int{0,1,0,3,12}
	fmt.Println("before: ", input)
	moveZeroes(input)
	fmt.Println("after: ", input)
}

// moveZeroes moves all zeros to the end of the array while keeping the order of the non-zeroes
//Input: [0,1,0,3,12]
//Output: [1,3,12,0,0]
func moveZeroes(nums []int)  {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], (append(nums[i+1:], nums[i]))...)
		}
	}
}
