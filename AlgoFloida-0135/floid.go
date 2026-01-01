package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	place := make([][]int, n)
	for i := 0; i < n; i++ {
		place[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&place[i][j])
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if place[i][k]+place[k][j] < place[i][j] {
					place[i][j] = place[i][k] + place[k][j]
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(place[i][j])
		}
		fmt.Println()
	}

}
