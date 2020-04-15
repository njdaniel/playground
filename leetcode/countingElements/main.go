package main

func main() {
	
}

// countElements counts how many consecutive numbers are +1 from the last
func countElements(arr []int) int {
	count := 0
	if len(arr) == 1 {
		return 0
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1]+1 {
			count++
		}
	}

	return count
}
