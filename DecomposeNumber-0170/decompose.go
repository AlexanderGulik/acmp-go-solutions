package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	result := 1
	for m := 2; m*(m+1) <= 2*n; m++ {
		num := 2*n - m*(m-1)
		if num > 0 && num%(2*m) == 0 {
			k := num / (2 * m)
			if k > 0 {
				result = m
			}
		}
	}

	for m := 1; m*(m+1) <= 2*n; m++ {
		if 2*n == m*(m+1) {
			result = m
			break
		}
	}
	fmt.Println(result)
}
