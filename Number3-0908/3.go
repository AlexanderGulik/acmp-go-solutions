package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)
	dp := make([]int, n+1)
	dp[1] = 0

	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + 1
		if i%2 == 0 {
			dp[i] = min(dp[i], dp[i/2]+1)
		}
		if i%3 == 0 {
			dp[i] = min(dp[i], dp[i/3]+1)
		}
	}
	fmt.Println(dp[n])
}
