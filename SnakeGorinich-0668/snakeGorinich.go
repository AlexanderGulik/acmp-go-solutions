package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	result := 0
	if m == 0 {
		if n%2 != 0 {
			fmt.Println(-1)
			return
		}
		result = n / 2
		fmt.Println(result)
		return
	}

	if n%2 == 0 {
		for m%4 != 0 {
			m = m - 1 + 2
			result++
		}
	} else {
		for m%4 != 2 {
			m = m - 1 + 2
			result++
		}
	}
	for m > 0 {
		m -= 2
		n += 1
		result++
	}
	if n%2 != 0 {
		fmt.Println(-1)
		return
	}

	result += n / 2

	fmt.Println(result)
}
