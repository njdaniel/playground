package main

import "fmt"

func main() {
	fmt.Println(romanToInt("LVIII"))
}

func romanToInt(s string) int {
	converts := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	runes := []rune(s)
	sum := 0
	var last rune
	for _, j := range runes {
		sum += converts[string(j)]
		switch {
		case (j == 'V' || j == 'X') && last == 'I':
			sum -= 2
		case (j == 'L' || j == 'C') && last == 'X':
			sum -= 20
		case (j == 'D' || j == 'M') && last == 'C':
			sum -= 200
		}
		last = j
	}
	return sum
}
