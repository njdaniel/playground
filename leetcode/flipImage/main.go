package main

import "fmt"

func main() {
	input := make([][]int, 0)
	input = [][]int{
		[]int{1,1,0,0,},[]int{1,0,0,1,},[]int{0,1,1,1,},[]int{1,0,1,0,},
	}
	fmt.Println(flipAndInvertImage(input))
}


func flipAndInvertImage(A [][]int) [][]int {
	 for _, arr := range A {
	 	for i := 0; i < len(arr)/2; i++ {
			tmp := 0
			tmp = arr[i]
			arr[i] = arr[len(arr)-1-i]
			arr[len(arr)-1-i] = tmp
		}
		for _, i := range arr {
			if i == 0 {
				i = 1
			} else if i == 1 {
				i = 0
			}
		}
	 }
	 return A
}