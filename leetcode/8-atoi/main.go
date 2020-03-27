package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	n := "-20000000000000000000"
	fmt.Println(myAtoi(n))
}

func myAtoi(str string) int {
	const MaxUint = ^uint32(0)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1
	str = strings.TrimSpace(str)
	str = strings.Split(str, " ")[0]
	str = strings.Split(str, ".")[0]
	i := 0
	if len(str) == 0 {
		return 0
	}

	if str[0] == '-' || str[0] == '+' {
		i = 1
	}
	for ;i < len(str); i++{
		if !unicode.IsDigit(rune(str[i])) {
			r := []rune(str)[:i]
			str = string(r)
			break
		}
	}
	if len(str) > 9 && str[0] != '-' {
		return MaxInt
	} else if len(str) > 9 && str[0] == '-' {
		return MinInt
	}
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	if n > MaxInt {
		n = MaxInt
	} else if n < MinInt {
		n = MinInt
	}
	return n
}