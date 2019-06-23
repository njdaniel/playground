package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := []int{0,1,2,4,5,7}
	fmt.Println(summaryRanges(input))

}

func summaryRanges(nums []int) []string {
	sr := make([]string, 0)
	c := make([]int, 0)

	if len(nums) == 0 {
		return sr
	}

	c = append(c, nums[0])

	for i := 1; i < len(nums); i++ {
		if c[len(c)-1]+1 == nums[i] {
			c = append(c, nums[i])
		} else {
			s := intToStringRange(c)
			sr = append(sr, s)
			c = c[:0]
			c = append(c, nums[i])
		}

		//if !(len(c) > 0) {
		//	c = append(c, i)
		//	continue
		//}
		//if i == c[len(c)-1] +1 {
		//	c = append(c, i)
		//	continue
		//} else {
		//	// create range string
		//	s := intToStringRange(c)
		//	sr = append(sr, s)
		//	// clear c
		//	c = c[:0]
		//}
		//switch {
		//case i == 0: c = append(c, nums[0])
		//case i == c[len(c)-1]+1: c = append(c, nums[i])
		//default:
		//	s := intToStringRange(c)
		//	sr = append(sr, s)
		//	c = c[:0]
		//}
	}
	if len(c) != 0 {
		s := intToStringRange(c)
		sr = append(sr, s)
	}
	return sr
}

func intToStringRange(nums []int) string {
	s := ""
	switch {
	case len(nums) == 0: return s
	case len(nums) == 1: s = strconv.Itoa(nums[0])
	case len(nums) > 1: s =  strconv.Itoa(nums[0]) + "->" + strconv.Itoa(nums[len(nums)-1])
	}
	return s
}
