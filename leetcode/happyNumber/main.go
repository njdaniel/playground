package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	input := 19
	fmt.Println(isHappy(input))
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
		buf := make([]byte, 1)
		_ = utf8.EncodeRune(buf, n[0])
		v, _ := strconv.Atoi(string(buf))
		if v == 1 {
			return v
		}
	}
	for _, v := range n {
		buf := make([]byte, 1)
		_ = utf8.EncodeRune(buf, v)
		v, _ := strconv.Atoi(string(buf))
		sum += v*v
	}

	sumr := make([]rune, 0)
	sumr = []rune(strconv.Itoa(sum))
	return squareDigits(sumr)

}