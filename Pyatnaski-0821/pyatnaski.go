package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var moves string
	fmt.Scan(&moves)
	count := 1
	table := make([][]int, n)
	for i := range table {
		table[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			table[i][j] = count
			count++
		}
	}
	table[n-1][n-1] = 0
	x, y := n-1, n-1
	runes := []rune(moves)
	for i := 0; i < len(moves); i++ {
		curr := runes[i]

		if curr == 'D' {
			if x+1 < n {
				table[x][y], table[x+1][y] = table[x+1][y], table[x][y]
				x += 1
			} else {
				fmt.Println("ERROR", i+1)
				return
			}
		} else if curr == 'L' {
			if y-1 >= 0 {
				table[x][y], table[x][y-1] = table[x][y-1], table[x][y]
				y -= 1
			} else {
				fmt.Println("ERROR", i+1)
				return
			}
		} else if curr == 'R' {
			if y+1 < n {
				table[x][y], table[x][y+1] = table[x][y+1], table[x][y]
				y += 1
			} else {
				fmt.Println("ERROR", i+1)
				return
			}
		} else if curr == 'U' {
			if x-1 >= 0 {
				table[x][y], table[x-1][y] = table[x-1][y], table[x][y]
				x -= 1
			} else {
				fmt.Println("ERROR", i+1)
				return
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", table[i][j])
		}
		fmt.Printf("\n")
	}
}
