package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	agents := make([][2]int, n)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		agents[i] = [2]int{a, b}
	}
	sort.Slice(agents, func(i, j int) bool {
		return agents[i][0] < agents[j][0]
	})

	dp := make([]int, n)
	if n >= 1 {
		dp[0] = 0
	}

	if n >= 2 {
		dp[1] = agents[1][1]
	}
	if n >= 3 {
		dp[2] = agents[1][1] + agents[2][1]
	}

	for i := 3; i < n; i++ {
		dp[i] = min(dp[i-1]+agents[i][1], dp[i-2]+agents[i][1])
	}
	fmt.Println(dp[n-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
