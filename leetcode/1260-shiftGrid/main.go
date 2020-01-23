package main

func main() {

}

func shiftGrid(grid [][]int, k int) [][]int {
	tmp := [][]int{}
	// [0][0]tmp = [len(grid)-1][[len(len(grid)-1][]grid)-1]grid
	m := len(grid) - 1
	n := len(grid[m]) - 1
	lastElem := grid[m][n]
	row0 := []int{lastElem}
	tmp = append(tmp, row0)
	for y, j := range grid {
		tmpRow := []int{}
		if y != 0 {
			tmpRow = append(tmpRow, grid[y-1][n])
		} else {
			tmpRow = row0
		}
		for x := 1; x < len(j); x++ {
			tmpRow = append(tmpRow, grid[y][x-1])
		}
		tmp = append(tmp, tmpRow)
	}
	return tmp
}
