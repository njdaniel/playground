package main

func main() {
	grid :=
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
	//tmp := make([]int, 0)
	tmp := s[len(s)-k:]
	s = s[:len(s)-k]
	s = append(tmp, s...)
	newGrid := make([][]int, 0)
	for x := 0; x < n; x++ {
		tmp := s[:m]
		newGrid = append(newGrid, tmp)
		s = s[m:]
	}
	return newGrid
}
