package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, k+2)
	}
	dp[0][n] = 1
	for i := 1; i <= k; i++ {
		for j := 1; j <= k+1; j++ {
			if j > 1 {
				dp[i][j-1] += dp[i-1][j]
			} else if j == 1 {
				if i == k {
					dp[i][0] += dp[i-1][j]
				}
			}
			if j < k+1 {
				dp[i][j+1] += dp[i-1][j]
			}
		}
	}
	fmt.Println(dp[k][0])
}
