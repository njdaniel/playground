package main

import "strconv"

func main() {
	
}

//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//1 is read off as "one 1" or 11.
//11 is read off as "two 1s" or 21.
//21 is read off as "one 2, then one 1" or 1211.
// 1 <= n <= 30
//not working
func countAndSay(n int) string {
	//recursive count until n is met
	si := []int{1}
	f := func() func() []int {
		r := 1
		count := 0
		tmp := make([]int, 0)
		return func() []int {
			for _, v := range si{
				if r == v {
					count++
				} else {
					tmp = append(tmp, count, r)
					r = v
				}
			}
			return tmp
		}
	}
	closure := f()
	sitmp := make([]int, 0)
	for i := 1; i >= n; i++ {
		sitmp = closure()
		n--
	}
	return ""
}