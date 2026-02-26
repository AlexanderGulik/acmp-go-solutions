package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	line, _ = reader.ReadString('\n')
	numStrs := strings.Fields(line)

	if n == 0 {
		fmt.Println(0)
		return
	}

	if n == 1 {
		fmt.Println(1)
		return
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(numStrs[i])
	}

	maxLen := 1
	currentLen := 1

	for i := 1; i < n; i++ {
		if arr[i] == arr[i-1] {

			if currentLen > maxLen {
				maxLen = currentLen
			}
			currentLen = 1
			continue
		}

		if currentLen == 1 {

			currentLen = 2
		} else {

			prevDiff := arr[i-1] - arr[i-2]
			currentDiff := arr[i] - arr[i-1]

			if (prevDiff > 0 && currentDiff < 0) || (prevDiff < 0 && currentDiff > 0) {
				currentLen++
			} else {

				if currentLen > maxLen {
					maxLen = currentLen
				}

				currentLen = 2
			}
		}
	}

	if currentLen > maxLen {
		maxLen = currentLen
	}

	fmt.Println(maxLen)
}
