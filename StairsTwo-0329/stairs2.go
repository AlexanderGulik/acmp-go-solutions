package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	dp := make([]int, n)
	prev := make([]int, n)
	dp[0] = arr[0]
	prev[0] = -1

	if n > 1 {
		if arr[1] > arr[0]+arr[1] {
			dp[1] = arr[1]
			prev[1] = -1
		} else {
			dp[1] = arr[0] + arr[1]
			prev[1] = 0
		}
	}
	for i := 2; i < n; i++ {
		op1 := dp[i-1] + arr[i]
		op2 := dp[i-2] + arr[i]
		if op1 >= op2 {
			dp[i] = op1
			prev[i] = i - 1
		} else {
			dp[i] = op2
			prev[i] = i - 2
		}
	}
	result := dp[n-1]
	fmt.Println(result)
	path := make([]int, 0)
	curr := n - 1
	for curr >= 0 {
		path = append([]int{curr + 1}, path...)
		curr = prev[curr]
	}
	for _, step := range path {
		fmt.Printf("%d ", step)
	}
}
