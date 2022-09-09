package main

func main() {

}

// generate creates a pascal triangle given the number of rows
//constraint: 1<=numRows<=30
func generate(numRows int) [][]int {
	rows := make([][]int, 0)
	if numRows < 1 {
		return rows
	}
	rows = append(rows, []int{1})

	for i := 1; i < numRows; i++ {
		newRow := make([]int, 0)
		newRow = append(newRow, 1)
		for j := 1; j < len(rows[i-1]); j++ {
			genNum := rows[i][j-1] + rows[i][j]
			newRow = append(newRow, genNum)
		}
		newRow = append(newRow, 1)
		rows = append(rows, newRow)
	}

	return rows
}
