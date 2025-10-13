package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	stations := n + 1
	dp := make([]int, stations)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i < stations-1; i++ {
		for j := i + 1; j < stations; j++ {
			var cost int
			fmt.Scan(&cost)
			if dp[i]+cost < dp[j] {
				dp[j] = dp[i] + cost
			}
		}
	}
	fmt.Println(dp[n])
}
