package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a--
		b--
		rootA := findRoot(a, parent)
		rootB := findRoot(b, parent)
		if rootA == rootB {
			continue
		}
		maxNum, minNum := max(rootA, rootB)
		arr[maxNum] += arr[minNum]
		arr[minNum] = 0
		parent[minNum] = maxNum

	}

	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			fmt.Printf("%d %d\n", i+1, arr[i])
		}
	}
}
func findRoot(x int, parent []int) int {
	for parent[x] != x {
		x = parent[x]
	}
	return x
}

func max(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

/*package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		maxNum, minNum := max(a, b)
		curr := arr[maxNum-1] + arr[minNum-1]
		arr[maxNum-1] = curr
		arr[minNum-1] = 0
	}
	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			fmt.Printf("%d %d\n", i+1, arr[i])
		}
	}
}

func max(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}*/
