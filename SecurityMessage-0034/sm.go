package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var str string
	fmt.Scan(&str)
	seen := make(map[string]int)
	for i := 0; i <= n-k; i++ {
		sub := str[i : i+k]
		seen[sub]++
		if seen[sub] >= 2 {
			fmt.Println("YES")
		}
	}
	fmt.Println("NO")
}
