package main

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
func countAndSay(n int) string {
	//recursive count until n is met
	f := func() {

	}
	for i := 1; i >= n; i++ {
		f()
		n--
	}
}