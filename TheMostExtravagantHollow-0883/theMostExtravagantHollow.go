package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scan(&n)
	var x, y float64

	for n > 0 {
		for x < 0.5 && n > 0 {
			fmt.Printf("%.6f %.6f\n", x, y)
			n--
			x = x + 0.01
		}
		x = 0
		y = y + 0.01
	}
}
