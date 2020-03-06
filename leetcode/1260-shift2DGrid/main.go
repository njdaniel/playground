package main

import "fmt"

func main() {
	grid := [][]int{{1}, {2}, {3}, {4}, {7}, {6}, {5}}
	k := 23
	fmt.Println(shiftGrid(grid, k))
}

func shiftGrid(grid [][]int, k int) [][]int {
	// plan:
	// data struct []int
	// convert grid to []int
	// put end of []int k values to temp
	// put at beginning
	// convert back to [][]int grid
	m := len(grid)
	var n int
	s := make([]int, 0)
	for _, row := range grid {
		n = len(row)
		s = append(s, row...)
	}
	fmt.Println("row lengh: ", n)
	totalLength := m*n
	if k > totalLength {
		k = k % totalLength
	} else if k % totalLength == 0 {
		return grid
	}

	//tmp := make([]int, 0)
	tmp := s[len(s)-k:]
	s = s[:len(s)-k]
	s = append(tmp, s...)
	newGrid := make([][]int, 0)
	for x := 0; x < m; x++ {
		tmp := s[:n]
		newGrid = append(newGrid, tmp)
		s = s[n:]
	}
	return newGrid
}
