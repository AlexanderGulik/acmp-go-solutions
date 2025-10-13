package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &numbers[i])
	}
	count, num := maxNumber(numbers)
	result := make([]int, 0, n)
	for i := 0; i < len(numbers); i++ {
		if numbers[i] != num {
			result = append(result, numbers[i])
		}
	}
	for i := 0; i < count; i++ {
		result = append(result, num)
	}
	for i := 0; i < len(result); i++ {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, result[i])
	}

}

func maxNumber(nums []int) (int, int) {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	counts := 0
	result := nums[0]
	for i, val := range count {
		if counts < val {
			counts = val
			result = i
		} else if counts == val && i < result {
			counts = val
			result = i
		}
	}
	return counts, result
}
