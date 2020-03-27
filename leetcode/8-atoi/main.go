package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	n := "4193 with words"
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
	//i := 0
	if len(str) == 0 {
		return 0
	}

	//if str[0] == '-' || str[0] == '+' {
	//	i = 1
	//}
	if len(str) == 0 {
		return 0
	}
	//get rid of leading zeros
	positive := true
	if str[0] == '-' {
		positive = false
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	for i := 0;i < len(str); i++{
		if !unicode.IsDigit(rune(str[i])) {
			r := []rune(str)[:i]
			str = string(r)
			break
		}
	}
	rev := ReverseString(str)
	for i := len(rev)-1; i > 0; i-- {
		if rev[i] != '0' {
			rev = rev[:i+1]
			break
		}
	}
	str = ReverseString(rev)

	if len(str) > 9 && positive{
		return MaxInt
	} else if len(str) > 9 && !positive {
		return MinInt
	}
	if !positive {
		str = fmt.Sprintf("-%s", str)
	}
	n, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	//if n > MaxInt {
	//	n = MaxInt
	//} else if n < MinInt {
	//	n = MinInt
	//}
	return n
}

func ReverseString(s string) string {
	rs := []rune(s)
	for i , j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}