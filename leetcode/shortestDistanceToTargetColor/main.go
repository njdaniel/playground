package main

import "fmt"

func main() {

	c := []int{1, 2, 3, 1}
	q := [][]int{{2, 1}, {2, 2}}
	fmt.Println(shortestDistanceColor(c, q))
}

//TODO: autocomplete comments like in VScode
//TODO: get info on Go key words
//TODO: go make key word using 0 vs specified value

//shortestDistanceColor
// given an array and query of index and value.
// return distance of value from index, if none return -1
func shortestDistanceColor(colors []int, queries [][]int) []int {
	r := make([]int, 0)
	//data structure
	// brain storm:
	//  1) linked list
	//  2) hash
	//     map[value int]  indexes []int | ordered?
	//     loop through colors to create map
	//     check if the value exist
	//     loop through indexes to find the lowest difference
	//
	//    optionb) only 3 values not too much sense to hash it.
	//
	//     O(1*ln(M) + N)
	//  3) list
	//     brute force loop forward and backward until match
	//     O(N)
	//     Probably going with this

	for _, v := range queries {
		shortest := -1
		func() {
			for i := v[0]; i < len(colors); i++ {

				if v[1] == colors[i] {
					distance := i - v[0]
					shortest = distance
				}
			}
			for i := v[0]; i >= 0; i-- {
				distance := v[0] - i
				if shortest != -1 {
					if distance < shortest {
						shortest = distance
					}
				} else {
					shortest = distance
				}
			}
		}()
		r = append(r, shortest)
	}

	return r
}
