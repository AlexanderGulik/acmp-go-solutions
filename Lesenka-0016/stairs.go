package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 3 {
		fmt.Println(1)
		return
	}
	memo := make(map[[2]int]int)
	var dfs func(level, n int) int

	dfs = func(level, n int) int {
		if n == 0 {
			return 1
		}
		key := [2]int{level, n}
		if val, ok := memo[key]; ok {
			return val
		}
		count := 0
		for i := 1; i < level && i <= n; i++ {
			count += dfs(i, n-i)
		}
		memo[key] = count
		return count
	}

	result := dfs(n+1, n)

	fmt.Println(result)
}

/*
func GetCount(level, n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for i := 1; i < level; i++ {
		count += GetCount(i, n-i)
	}
	return count
}
*/
