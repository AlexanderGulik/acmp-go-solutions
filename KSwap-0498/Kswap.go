package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	used := make([]bool, n+1)
	count := 0
	var bactrack func(prev int, depth int)
	bactrack = func(prev int, depth int) {
		if depth == n {
			count++
			return
		}
		for i := 1; i <= n; i++ {
			if !used[i] {
				if depth == 0 || abs(prev-i) <= k {
					used[i] = true
					bactrack(i, depth+1)
					used[i] = false
				}
			}
		}
	}

	bactrack(0, 0)
	fmt.Println(count)

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
