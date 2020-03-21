package main

import (
	"fmt"
	"strconv"
)

func main() {
	var r int64
	r = 2
	arr := make([]int64, 0)
	arr = []int64{1, 2, 2, 4}
	fmt.Println(countTriplets(arr, r))
}

// Complete the countTriplets function below.
// find num of triplets with common ratio r and i<j<k
// assume arr passed in order
// assume len(arr) > 2
func countTriplets(arr []int64, r int64) int64 {
	// store the positions of each triplet
	// check if used and store to return num of triplets
	trips := make(map[string]interface{},0)
	// N(O^3) for loop for each
	for i := 0; i < len(arr) - 2; i++ {
		for j := 1; j < len(arr) -1; j++ {
			for k :=2; k < len(arr); k++ {
				pos := make([]rune, 3)
				pos[0] = rune(strconv.Itoa(i)[0])
				pos[1] = rune(strconv.Itoa(j)[0])
				pos[2] = rune(strconv.Itoa(k)[0])
				if _, ok := trips[string(pos)]; !ok {
					if arr[i]*r==arr[j] && arr[j]*r==arr[k] {
						trips[string(pos)] = nil
					}
				}
			}
		}
	}
	return int64(len(trips))
}