package main

func main() {
	
}

func maxSubArray(nums []int) int {
	var f *int
	var e *int
	//f = &nums[0]
	//e = &nums[0]
	for x := 0; x < len(nums); x++ {
		if f == nil {
			f = &nums[x]
		}
		if e == nil {
			e = &nums[x]
		}

		//should end move up?

		//should front move up?
		v := 0
		for y := *f; y <= x; y++ {
			v += nums[y]
		}
		if v > 0 {
			f = &nums[x]
		}
	}
}
