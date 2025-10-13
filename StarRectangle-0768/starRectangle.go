package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				xx, yy := 0, 0
				if check(x, y, i, j, k, &xx, &yy) {
					for ii := k + 1; ii < n; ii++ {
						if x[ii] == xx && y[ii] == yy {
							res++
						}
					}
				}
			}
		}
	}
	fmt.Println(res)
}

func check(x, y []int, i, j, k int, xx, yy *int) bool {
	if (x[j]-x[i])*(x[k]-x[i])+(y[j]-y[i])*(y[k]-y[i]) == 0 {
		*xx = x[j] - x[i] + x[k]
		*yy = y[j] - y[i] + y[k]
		return true
	}
	if (x[i]-x[j])*(x[k]-x[j])+(y[i]-y[j])*(y[k]-y[j]) == 0 {
		*xx = x[i] - x[j] + x[k]
		*yy = y[i] - y[j] + y[k]
		return true
	}
	if (x[j]-x[k])*(x[i]-x[k])+(y[j]-y[k])*(y[i]-y[k]) == 0 {
		*xx = x[j] - x[k] + x[i]
		*yy = y[j] - y[k] + y[i]
		return true
	}
	return false
}

/*package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func main() {
	var n int
	fmt.Scan(&n)

	if n < 4 {
		fmt.Println(0)
		return
	}

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		points[i] = Point{x, y}
	}

	result := countRectangles(points)
	fmt.Println(result)
}

func countRectangles(points []Point) int {
	count := 0
	pointSet := make(map[Point]bool)
	for _, p := range points {
		pointSet[p] = true
	}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]
			p3 := Point{p1.x, p2.y}
			p4 := Point{p2.x, p1.y}
			if p3 != p1 && p3 != p2 && p4 != p1 && p4 != p2 &&
				pointSet[p3] && pointSet[p4] {
				count++
			}
		}
	}
	return count / 2
}*/
