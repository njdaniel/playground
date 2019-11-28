package main

func main() {
	
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
			switch arr[i] {
			case 0:
				arr[i] = 1
			case 1:
				arr[i] = 0
			}
		}
	 }
	 return A
}