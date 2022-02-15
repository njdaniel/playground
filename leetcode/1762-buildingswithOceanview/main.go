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

	return buildingswithOceanview
}
