package main

import "fmt"

func main() {
	fmt.Println(rotLeft([]int32{1, 2, 3}, 2))
}

func rotLeft(a []int32, d int32) []int32 {

	result := make([]int32, len(a))
	for k, v := range a {
		newk := k - int(d)
		if newk < 0 {
			newk = len(a) - int(d) + k
		}
		result[newk] = v
	}
	return result
}
