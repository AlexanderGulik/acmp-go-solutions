package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	if n == 1 {
		fmt.Println(0)
		return
	}
	dp := make([]int, n)
	dp[0] = 0
	if n > 1 {
		dp[1] = mod(nums[1] - nums[0])
	}
	for i := 2; i < n; i++ {
		jump1 := dp[i-1] + mod(nums[i]-nums[i-1])
		jump2 := dp[i-2] + 3*mod(nums[i]-nums[i-2])
		dp[i] = min(jump2, jump1)
	}
	fmt.Println(dp[n-1])

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func mod(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
