package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := "-91283472332"
	fmt.Println(myAtoi(n))
}

func myAtoi(str string) int {
	const MaxUint = ^uint32(0)
	fmt.Println(MaxUint)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1
	str = strings.TrimSpace(str)
	str = strings.Split(str, " ")[0]
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