package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var a, b, k int
	fmt.Scan(&a, &b, &k)
	na := (a - 1) / k
	nb := (b - 1) / k
	if nb-na >= n-1 {
		nb = na + n - 1
	}
	m := -1
	for steps := na; steps <= nb; steps++ {
		pos1 := steps % n
		if arr[pos1] > m {
			m = arr[pos1]
		}
		pos2 := (n - steps%n) % n
		if arr[pos2] > m {
			m = arr[pos2]
		}
	}
	fmt.Println(m)
}
