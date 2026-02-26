package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	coords := make([][10]float64, n)
	for i := 0; i < n; i++ {
		for j := 0; j < 10; j++ {
			fmt.Scan(&coords[i][j])
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		x, y := coords[i][0], coords[i][1]
		x1, y1 := coords[i][2], coords[i][3]
		x2, y2 := coords[i][4], coords[i][5]
		x3, y3 := coords[i][6], coords[i][7]
		x4, y4 := coords[i][8], coords[i][9]

		if pointInRectangle(x, y, x1, y1, x2, y2, x3, y3, x4, y4) {
			result++
		}
	}
	fmt.Println(result)
}
func pointInRectangle(px, py, x1, y1, x2, y2, x3, y3, x4, y4 float64) bool {
	cross := func(ax, ay, bx, by float64) float64 {
		return ax*by - ay*bx
	}

	d1 := cross(x2-x1, y2-y1, px-x1, py-y1)
	d2 := cross(x3-x2, y3-y2, px-x2, py-y2)
	d3 := cross(x4-x3, y4-y3, px-x3, py-y3)
	d4 := cross(x1-x4, y1-y4, px-x4, py-y4)

	if (d1 >= 0 && d2 >= 0 && d3 >= 0 && d4 >= 0) ||
		(d1 <= 0 && d2 <= 0 && d3 <= 0 && d4 <= 0) {
		return true
	}
	return false
}
