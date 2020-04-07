package bruteforce

// MaxProfit returns the highest profit
func MaxProfit(prices []int) int {
	return calc(prices, 0)
}

func calc(prices []int, s int) int {
	if s >= len(prices) {
		return 0
	}
	max := 0
	for start := s; start < len(prices); start++ {
		maxprofit := 0
		for i := start+1; i < len(prices); i++ {
			if prices[start] <  prices[i] {
				profit := calc(prices, i + 1) + prices[i] - prices[start]
				if profit > maxprofit {
					maxprofit = profit
				}
			}
		}
		if maxprofit > max {
			max = maxprofit
		}
	}
	return max
}
