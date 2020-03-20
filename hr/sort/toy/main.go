package main

import (
	"fmt"
	"sort"
)

func main() {
	prices := []int32{1, 12, 5, 10}
	saving := int32(4)
	fmt.Println(maximumToys(prices, saving))
}

type lowHi []int32

func (l lowHi) Len() int {
	return len(l)
}

func (l lowHi) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l lowHi) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Complete the maximumToys function below.
func maximumToys(prices []int32, k int32) int32 {
	counter := 0
	//sort prices by lowest to highest
	sort.Sort(lowHi(prices))
	//loop through and subtract from k(budget), until negative(break), zero(+1, break), positive(cont)
	//kint := int(k)
	for _, v := range prices {
		k = k-v
		switch {
		case k > int32(0):
			counter++
		case k == int32(0):
			counter++
			break
		case k < int32(0):
			break
		}
	}
	return int32(counter)
}
