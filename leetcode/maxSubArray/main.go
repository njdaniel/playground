package main

import "fmt"

func main() {
	input := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(input))
	
}

func maxSubArray(nums []int) int {
	var f int
	var e int
	//f = &nums[0]
	//e = &nums[0]
	for x := 0; x < len(nums); x++ {
		//if f == nil && e == nil {
		//	f = x
		//	e = x
		//	fmt.Println("init")
		//	continue
		//}


		//should front move up?
		vf := 0
		for y := f; y <= x; y++ {
			vf += nums[y]
		}
		if vf > 0 {
			f = x
		}

		//should end move up?
		ve := 0
		for y := e; y <= x || y <= f; y++ {
			ve += nums[y]
		}
		if ve < 0 {
			e = x
		}
		sum := 0
		for y := e; y <= f; y++ {
			sum += nums[y]
		}
		fmt.Printf("front pointer %d\n", f)
		fmt.Printf("end pointer %d\n", e)
		fmt.Printf("round %d sum: %d\n\n", x, sum)
	}

	sum := 0
	for y := e; y <= f; y++ {
		sum += nums[y]
	}
	return sum
}
