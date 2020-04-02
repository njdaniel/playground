package main

import "strconv"

func main() {
	
}

func isHappy(n int) bool {
	nr := make([]rune, 0)
	nr = []rune(strconv.Itoa(n))
	n = squareDigits(nr)
	if n == 1 {
		return true
	}
	return false
}

func squareDigits(n []rune) int {
	sum := 0
	if len(n) == 1 {
		return sum
	}
	for _, v := range n {
		sum += int(v)*int(v)
	}

	sumr := make([]rune, 0)
	sumr = []rune(strconv.Itoa(sum))
	squareDigits(sumr)
	if len(sumr) == 1 {
		return sum
	}
	return 0

}