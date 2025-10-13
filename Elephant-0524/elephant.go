package main

import "fmt"

func main() {
	var m, n, x, y int
	fmt.Scan(&m, &n, &x, &y)
	if m == x && n == y {
		fmt.Println(0)
		return
	}
	if (m+n)%2 != (x+y)%2 {
		fmt.Println(0)
		return
	}

	dx := abs(m - x)
	dy := abs(n - y)
	if dx == dy {
		fmt.Println(1)
		return
	}

	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			if abs(i-m) == abs(j-n) && abs(i-x) == abs(j-y) {
				fmt.Println(2)
				fmt.Println(i, j)
				return
			}
		}
	}

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
