package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	edges := make([][3]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&edges[i][0], &edges[i][1], &edges[i][2])
	}

	dist := make([]int, n)

	for i := 0; i < n; i++ {
		dist[i] = 30000
	}
	dist[0] = 0

	for i := 0; i < n-1; i++ {
		up := false
		for j := 0; j < m; j++ {
			from := edges[j][0] - 1
			to := edges[j][1] - 1
			weight := edges[j][2]

			if dist[from] < 30000 && dist[from]+weight < dist[to] {
				dist[to] = dist[from] + weight
				up = true
			}
		}

		if !up {
			break
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", dist[i])
	}
}
