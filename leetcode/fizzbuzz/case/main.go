package main

import (
	"fmt"
	"strconv"
)

func main() {
	fizzbuzz(15)
}
//fizzbuzz count up to given num(n)
// if div by 3 fizz
//  if div by 5 buzz
// print out each line
func fizzbuzz(n int)  {
	for i := 1; i <=n; i++ {
		ls := make([]rune, 0)
		switch {
		case i % 3 == 0:
			ls = append(ls, []rune("fizz")...)
			fallthrough
		case i % 5 == 0:
			if i % 5 == 0 {
				ls = append(ls, []rune("buzz")...)
			}
		default:
			ls = append(ls, []rune(strconv.Itoa(i))...)
		}
		fmt.Println(string(ls))
	}
}
