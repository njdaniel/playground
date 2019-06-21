package main

import "math"

func main() {

}

func maxSubArray(nums []int) int {
	curMax := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		curMax = int(math.Max(float64(nums[i]), float64(curMax+nums[i])))
		max = int(math.Max(float64(max), float64(curMax)))
	}
	return max
}