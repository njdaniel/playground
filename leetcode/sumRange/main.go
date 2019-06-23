package main

import "strconv"

func main() {
	
}

func summaryRanges(nums []int) []string {
	sr := []string{}
	c := []int{}

	for i := range nums {
		if !(len(c) > 0) {
			c = append(c, i)
			continue
		}
		if i == c[len(c)-1] +1 {
			c = append(c, i)
			continue
		} else {
			// create range string
			intToStringRange(c)
			// clear c
			c = c[:0]
		}
	}
	return sr
}

func intToStringRange(nums []int) string {
	s := ""
	switch {
	case len(nums) == 0: return s
	case len(nums) == 1: s = strconv.Itoa(nums[0])
	case len(nums) > 1: s =  strconv.Itoa(nums[0]) + " -> " + strconv.Itoa(nums[len(nums)-1])
	}
	return s
}
