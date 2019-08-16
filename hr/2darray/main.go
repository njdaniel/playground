package main

import "fmt"

func main() {
	
}


// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	// assume input is a 6x6
	//	hourglass is 313
	var sumMax int32
	var hg func(x, y int) int32
	fmt.Printf("intersect: %d\n", arr[0][1])
	hg = func(x, y int) int32 {
		sum := arr[x][y] + arr[x][y+1] + arr[x][y+2] + arr[x+1][y+1] + arr[x+2][y] + arr[x+2][y+1] + arr[x+2][y+2]
		return sum
	}
	for i :=0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := hg(i,j)
			if i ==0 && j == 0 {
				sumMax = sum
			} else {
				if sum > sumMax {
					sumMax = sum
				}
			}
		}
	}
	return sumMax
}