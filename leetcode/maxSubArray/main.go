package main

import "fmt"

func main() {
	input := []int{-2, -1}
	fmt.Println(maxSubArray(input))
	
}

func maxSubArray(nums []int) int {
	var f int
	var e int
	sum := nums[0]
	for x := 1; x < len(nums); x++ {
		//should front move up?
		vf := sum
		for y := f+1; y <= x; y++ {
			vf += nums[y]
		}
		if vf > sum {
			f = x
		}

		//should end move up?
		ve := sum
		for y := e; y < x && y <= f; y++ {
			ve += nums[y]
		}
		if ve < sum && e <= f && f >= x {
			e = x
		}
		if f==e && nums[x] > nums[f] {
			f = x
			e = x
		}
		newSum := 0
		for z := e; z <= f; z++ {
			newSum += nums[z]
		}
		sum = newSum
		fmt.Printf("front pointer %d\n", f)
		fmt.Printf("end pointer %d\n", e)
		fmt.Printf("round %d sum: %d\n\n", x, sum)
	}
	return sum
}
