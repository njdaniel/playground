package main

import "fmt"

func main() {
	
}

func minimumBrides(q []int32)  {
	bribes := 0
	for k, v := range q {
		if k+3 < int(v) {
			fmt.Println("Too chaotic")
			return
		}
		// for each value that's higher in front, increase the bribe
		// counting how much each value gets bribed
		for i := k-1; i >= 0; i-- {
			if q[i] > v {
				bribes++
			}
		}
	}

	fmt.Println(bribes)
 }