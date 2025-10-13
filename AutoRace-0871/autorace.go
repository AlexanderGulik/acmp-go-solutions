package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	remeber := make(map[[2]int]bool)
	graph := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		low, high := a, b
		if a > b {
			low, high = b, a
		}
		z := [2]int{low, high}
		if !remeber[z] {
			remeber[z] = true
			graph[a] = append(graph[a], b)
			graph[b] = append(graph[b], a)
		}
	}
	visited := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if !visited[i] {
			if hasCycle(graph, visited, i, -1) {
				fmt.Println("YES")
				return
			}
		}
	}
	fmt.Println("NO")

}

func hasCycle(graph [][]int, visited []bool, node, parent int) bool {
	visited[node] = true

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			if hasCycle(graph, visited, neighbor, node) {
				return true
			}
		} else if neighbor != parent {
			return true
		}
	}
	return false
}
