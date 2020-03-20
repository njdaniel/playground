package main

func main() {
	
}

// sherlockAndAnagrams
// for each repeated number, increase the number of x! pairs of characters
func sherlockAndAnagrams(s string) int32 {
	pairs := 0
	chars := make(map[rune]int, 0)
	for _, v := range s {
		if _, ok := chars[v]; !ok {
			chars[v] = 1
		} else {
			chars[v]++
			pairs++
		}
	}
	return int32(triangleNum(pairs))
}

func triangleNum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}