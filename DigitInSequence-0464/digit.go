package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	if n == 1 {
		fmt.Println(0)
		return
	}
	lenght := int64(1)
	for lenght < n {
		lenght *= 2
	}
	result := findDigit(n, lenght)
	fmt.Println(result)
}

func findDigit(pos, blockSize int64) int {
	if blockSize == 1 {
		return 0
	}
	half := blockSize / 2
	if pos <= half {
		return findDigit(pos, half)
	} else {
		prev := findDigit(pos-half, half)
		return (prev + 1) % 3
	}
}

/*
func main() {
	var n int
	fmt.Scan(&n)
	seq := []int{0}
	for len(seq) <= n {
		newPart := make([]int, len(seq))
		copy(newPart, seq)
		for i := 0; i < len(seq); i++ {
			if newPart[i] == 2 {
				newPart[i] = 0
			} else {
				newPart[i]++
			}
		}
		seq = append(seq, newPart...)
	}
	fmt.Println(seq[n-1])
}
*/
