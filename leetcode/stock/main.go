package main

import "fmt"

func main() {
	input := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(input))
}

func maxProfit(prices []int) int {
	// find the greatest difference between
	// loop through compare with rest
	profit := 0
	for k, v := range prices {
		for i := k; i < len(prices); i++ {
			if prices[i] - v > profit {
				profit = prices[i] - v
			}
		}
	}
	return profit
}
