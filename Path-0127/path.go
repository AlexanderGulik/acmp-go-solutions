package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&graph[i][j])
		}
	}
	var start, end int
	fmt.Scan(&start, &end)
	start--
	end--
	if start == end {
		fmt.Println(0)
		return
	}
	visit := make([]bool, n)
	dist := make([]int, n)
	queue := make([]int, 0)

	visit[start] = true
	dist[start] = 0
	queue = append(queue, start)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for neig := 0; neig < n; neig++ {
			if graph[curr][neig] == 1 && !visit[neig] {
				visit[neig] = true
				dist[neig] = dist[curr] + 1
				if neig == end {
					fmt.Println(dist[neig])
					return
				}
				queue = append(queue, neig)
			}
		}
	}
	fmt.Println(-1)
}
