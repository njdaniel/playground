package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addBinary("10", "01"))
}

func addBinary(a, b string) string {
	// ai, _ := strconv.Atoi(a)
	// bi, _ := strconv.Atoi(b)
	ab, _ := strconv.ParseInt(a, 2, 64)
	bb, _ := strconv.ParseInt(b, 2, 64)
	fmt.Println(ab)
	fmt.Println(bb)
	cb := ab + bb
	cs := strconv.FormatInt(cb, 2)
	return cs
}
