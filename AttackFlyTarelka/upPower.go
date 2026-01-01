// первая задача в 26 году, после сложных 60 дней
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var xp1, yp1, xp2, yp2 int
	fmt.Scan(&xp1, &yp1, &xp2, &yp2)
	result := []int{}
	a := yp2 - yp1
	b := xp1 - xp2
	c := -(a*xp1 + b*yp1)
	for i := 1; i <= n; i++ {
		//Как узнать что точка на кординатной плоскости x y, пересекает линию между x1 y1 x2 y2
		// забыл про радиус, задача пиздата
		// не прочитал про беск луч спс ребята,-1ч кайфуем на отрезочке))))
		// когда то явно я научсь читать, 2 тест не пройден. пиздец..

		var x, y, r int
		fmt.Scan(&x, &y, &r)
		num := math.Abs(float64(a*x + b*y + c))
		den := math.Sqrt(float64(a*a + b*b))

		var d float64
		if den == 0 {
			dx := x - xp1
			dy := y - yp1
			d = math.Sqrt(float64(dx*dx + dy*dy))
		} else {
			d = num / den
		}

		if d <= float64(r)+1e-9 {
			result = append(result, i)
		}

	}

	fmt.Printf("%d \n", len(result))
	for i := 0; i < len(result); i++ {
		fmt.Printf("%d ", result[i])
	}
}

/*
уйди из моей головы, я не могу не испытывать к тебе чувства
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
*/

/*if min(xp1, xp2) <= x <= max(xp1, xp2) && min(yp1, yp2) <= y <= max(yp1, yp2) {
	result = append(result, i)
}*/

/*
я еьал в рот условие с беск лучом

	for i := 0; i < n; i++ {
		//Как узнать что точка на кординатной плоскости x y, пересекает линию между x1 y1 x2 y2
		// забыл про радиус, задача становистя сложнее
		var x, y, r int
		fmt.Scan(&x, &y, &r)

		vx := xp2 - xp1
		vy := yp2 - yp1

		if vx == 0 && vy == 0 {
			dx := x - xp1
			dy := y - yp1
			if dx*dx+dy*dy <= r*r {
				result = append(result, i)
			}
			continue
		}
		wx := x - xp1
		wy := y - yp1

		dot := wx*vx + wy*vy
		lenSq := vx*vx + vy*vy
		t := float64(dot) / float64(lenSq)
		if t < 0 {
			t = 0
		} else if t > 1 {
			t = 1
		}
		closetx := float64(xp1) + t*float64(vx)
		closety := float64(yp1) + t*float64(vy)

		dx := float64(x) - closetx
		dy := float64(y) - closety
		distSq := dx*dx + dy*dy

		if distSq <= float64(r*r)+1e-9 {
			result = append(result, i)
		}

		/*if min(xp1, xp2) <= x <= max(xp1, xp2) && min(yp1, yp2) <= y <= max(yp1, yp2) {
			result = append(result, i)
		}*/
