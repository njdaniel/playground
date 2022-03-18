package main

func main() {

}

func longestPalindrome(s string) int {
	//edge cases:
	//length 0 1
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	//convert to []rune
	//count pairs by finding and removing from slice
	//return counted up pairs *2 plus 1 if there are spares remaining
	rs := []rune(s)
	pairs := 0
	remainder := 0
	for i := 0; i < len(rs); i++ {
		for j := i + 1; j < len(rs); j++ {
			if rs[i] == rs[j] {
				pairs++
				//remove pairs
				if j != len(rs)-1 {
					rs = append(rs[:j], rs[j+1:]...)
				} else {
					rs = rs[:len(rs)-1]
				}
				if i != len(rs)-1 {
					rs = append(rs[:i], rs[i+1:]...)
				} else {
					rs = rs[:len(rs)-1]
				}

				//reset loop
				i = -1
				break
			}
		}

	}
	if len(rs) > 0 {
		remainder = 1
	}

	return pairs*2 + remainder
}
