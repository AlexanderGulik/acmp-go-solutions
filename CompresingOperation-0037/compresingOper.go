package main

import (
	"fmt"
	"math"
)

type storage struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {
	var n int
	var q float64
	fmt.Scan(&n)
	fmt.Scan(&q)
	Dotes := make([]storage, n)
	for i := 0; i < n; i++ {
		var a, b, c, d int
		fmt.Scan(&a, &b, &c, &d)
		Dotes[i] = storage{a, b, c, d}
	}
	isCompressing := true
	for _, p := range Dotes {
		normX := math.Sqrt(float64(p.x1*p.x1 + p.y1*p.y1))
		normAX := math.Sqrt(float64(p.x2*p.x2 + p.y2*p.y2))
		if normAX > q*normX+1e-9 {
			isCompressing = false
			break
		}
	}
	if isCompressing {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
