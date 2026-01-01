package main

import "fmt"

func main() {
	var n, m int
	var k int
	fmt.Scan(&n, &m, &k)
	space := make([][]int, n)
	for i := 0; i < n; i++ {
		space[i] = make([]int, m)
	}

	for i := 0; i < k; i++ {
		var a, y, x int
		fmt.Scan(&a, &y, &x)
		x--
		y--
		if y < 0 || y+1 >= n || x < 0 || x+1 >= m {
			continue
		}
		var cells [][2]int

		if a == 1 {
			cells = [][2]int{{y, x + 1}, {y + 1, x}, {y + 1, x + 1}}
		}

		if a == 2 {
			cells = [][2]int{{y, x}, {y + 1, x}, {y + 1, x + 1}}
		}

		if a == 3 {
			cells = [][2]int{{y, x}, {y, x + 1}, {y + 1, x + 1}}
		}

		if a == 4 {
			cells = [][2]int{{y, x}, {y, x + 1}, {y + 1, x}}
		}
		canPlace := true
		for _, cell := range cells {
			r, c := cell[0], cell[1]
			if space[r][c] != 0 {
				canPlace = false
				break
			}
		}
		if canPlace {
			for _, cell := range cells {
				r, c := cell[0], cell[1]
				space[r][c] = 1
			}
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if space[i][j] == 1 {
				result++
			}
		}
	}
	fmt.Println(result)

}
