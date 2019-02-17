package main

import "fmt"

func main() {
	s := []byte{'1', '2', '3', '4', '5'}
	fmt.Println(len(s)/2)
	n1 := 1.8
	n2 := 1.8
	fmt.Println(int(n1)+int(n2))
}

func reverseSlice(a []byte)  {
	for i := len(a)/2-1; i >= 0; i-- {
		opp := len(a)-1-i
		a[i], a[opp] = a[opp], a[i]
	}
}
