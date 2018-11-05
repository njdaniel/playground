package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}

//climbStairs either 1 or 2 steps
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		// 2 ways 11 or 2
		return 2
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2 // 2 ways 11 or 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}
