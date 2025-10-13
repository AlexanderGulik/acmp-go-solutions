package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Ints(arr)
	if n == 2 {
		fmt.Println(arr[1] - arr[0])
		return
	}
	dp := make([]int, n)
	dp[0] = 0
	dp[1] = arr[1] - arr[0]
	if n >= 3 {
		dp[2] = arr[2] - arr[1] + dp[1]
	}
	for i := 3; i < n; i++ {
		dp[i] = min(dp[i-1]+arr[i]-arr[i-1], dp[i-2]+arr[i]-arr[i-1])
	}
	fmt.Println(dp[n-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
