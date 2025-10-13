package main

import "fmt"

const INF = 1 << 30

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, n+1)
		for j := range graph[i] {
			graph[i][j] = INF
		}
	}

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		if c < graph[a][b] {
			graph[a][b] = c
			graph[b][a] = c
		}
	}
	visited := make([]bool, n+1)
	minEdge := make([]int, n+1)
	for i := range minEdge {
		minEdge[i] = INF
	}
	minEdge[1] = 0
	totalWeight := 0
	for i := 1; i <= n; i++ {
		v := -1
		for j := 1; j <= n; j++ {
			if !visited[j] && (v == -1 || minEdge[j] < minEdge[v]) {
				v = j
			}
		}
		visited[v] = true
		totalWeight += minEdge[v]
		for j := 1; j <= n; j++ {
			if graph[v][j] < minEdge[j] && !visited[j] {
				minEdge[j] = graph[v][j]
			}
		}
	}

	fmt.Println(totalWeight)

}
