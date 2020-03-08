package main

import "fmt"

func main() {
	
}

func minimumBrides(q []int32)  {
	bribes := 0
	for k, v := range q {
		if k+3 >= int(v) {
			fmt.Println("Too chaotic")
			return
		}
		if k+1 > int(v) {
			bribes += ((k+1)-int(v))
		}
	}
	fmt.Println(bribes)
}