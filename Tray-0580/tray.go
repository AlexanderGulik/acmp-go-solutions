package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	var tr float64
	fmt.Scan(&a, &b, &c)
	fmt.Scan(&tr)
	p := (a + b + c) / 2
	S := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	r := S / p
	if r >= tr {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
