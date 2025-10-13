package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, size int
	fmt.Scan(&n, &size)
	queue := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&queue[i])
	}
	sort.Ints(queue)
	result := 0
	left, right := 0, len(queue)-1
	for left <= right {
		if queue[left]+queue[right] <= size {
			result++
			left++
			right--
		} else {
			result++
			right--
		}
	}
	fmt.Println(result)
}
