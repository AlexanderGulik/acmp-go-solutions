package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	product := 1.0
	for i := 0; i < n; i++ {
		product *= (2*arr[i] - 1)
	}
	result := (1 + product) / 2
	fmt.Printf("%.6f", result)

}
