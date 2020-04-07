package main

import "fmt"

func main() {
	input := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(input))
}

func maxProfit(prices []int) int {
	buy := true
	profit := 0
	stockValue := 0
	for i, price := range prices {
		//check if its the last day
		if i == len(prices)-1 {
			if !buy {
				if price > stockValue {
					profit += (price -stockValue)
				}
			}
			break
		}
		if buy {
			//check if should buy
			// 1. a higher price to sell at
			//
			for j := i+1; j < len(prices); j++ {
				if price < prices[j] {
					stockValue = price
					buy = false
					break
				}
			}
		} else {
			// check if should sell
			// 1. this is the highest price to sell at w/o the next day being higher OR its the last day
			if price > stockValue && prices[i] > prices[i+1] {
				profit += (price-stockValue)
				buy = true
			}

		}
	}
	return profit
}