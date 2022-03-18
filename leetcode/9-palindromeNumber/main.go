package main

import "strconv"

func main() {

}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		if rs[i] != rs[j] {
			return false
		}
	}
	return true
}

func isPalindromeNoString(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	for i, y := 0, 0; x > y; i++ {
		y = y*10 + (x % 10)
		if x == y || x/10 == y {
			return true
		}
		x = x / 10
	}
	return false
}
