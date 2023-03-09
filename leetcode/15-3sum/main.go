package main

import "errors"

func main() {

}

func threeSum(nums []int) [][]int {
	output := make([][]int, 0)
	n1 := 0
	n2 := 0
	n3 := 0

	for i := range nums {
		n1 = nums[i]
		for j := i + 1; j < len(nums); j++ {
			n2 = nums[j]
			if i == j {
				break
			}
			for k := j + 1; k < len(nums); k++ {
				n3 = nums[k]
				if j == k {
					break
				}
				if i == k {
					break
				}
				if n1+n2+n3 == 0 {
					success := []int{n1, n2, n3}
					successDupe := false
					if len(output) > 0 {
						for i := range output {
							successDupe = isDupe(success, output[i])

						}
					}
					if !successDupe {

						output = append(output, success)
					}
				}
			}
		}
	}
	// toRemove := []int{}
	// if len(output) > 1 {
	// 	for i := 0; i < len(output); i++ {
	// 		for j := 1; j < len(output); j++ {
	// 			if isDupe(output[i], output[j]) {
	// 				toRemove = append(toRemove, j)
	// 			}
	// 		}
	// 	}
	// }

	return output
}

func isDupe(nums1, nums2 []int) bool {
	d := false
	var s stack
	for _, v := range nums1 {
		s.Push(v)
	}
	for i := 0; i < len(nums2)^2; i++ {
		if nums2[i] == s[len(s)-1] {
			_, _ = s.Pop()
		}
	}
	if len(s) == 0 {
		d = true
	}
	return d
}

type stack []int

func (s *stack) Push(v int) stack {
	return append(*s, v)
}

func (s *stack) Pop() (int, error) {
	l := len(*s)
	if l == 0 {
		return 0, errors.New("empty stack")
	}
	res := (*s)[l-1]
	news := (*s)[:l-1]
	s = &news
	return res, nil
}
