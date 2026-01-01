package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	xx := math.Inf(-1)
	yy := math.Inf(-1)
	ss := math.Inf(1)
	x := make([]float64, n)
	y := make([]float64, n)
	s := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i], &s[i])
		xx = max(xx, x[i])
		yy = max(yy, y[i])
	}
	for i := 0; i < n; i++ {
		ss = min(ss, s[i]-(xx-x[i])-(yy-y[i]))
	}
	if ss > 0 {
		result := ss * ss / 2
		fmt.Printf("%.3f", result)
	} else {
		fmt.Println(0)
	}

}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
