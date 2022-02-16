package main

import "fmt"

func main() {
	heights := []int{4, 2, 3, 1}
	a := findBuildings(heights)
	fmt.Println(a)
}

func findBuildings(heights []int) []int {
	buildingswithOceanview := make([]int, 0)
	//bruteforce: loop through each input and compare to the rest of slice to the right
	//  if its > all then add to list and return
	for i, v := range heights {
		oceanView := true
		for j := i + 1; j < len(heights); j++ {
			if v <= heights[j] {
				oceanView = false
				break
			}
		}
		if oceanView {
			buildingswithOceanview = append(buildingswithOceanview, i)
		}
	}

	return buildingswithOceanview
}
