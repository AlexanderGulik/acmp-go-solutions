package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	line, _ := reader.ReadString('\n')
	nm := strings.Fields(line)
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	line, _ = reader.ReadString('\n')
	firstLine := strings.Fields(line)
	nums := make(map[int]bool)

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(firstLine[i])
		nums[num] = true
	}
	line, _ = reader.ReadString('\n')
	secondLine := strings.Fields(line)
	remember := make(map[int]bool)
	for i := 0; i < m; i++ {
		num, _ := strconv.Atoi(secondLine[i])
		if nums[num] {
			remember[num] = true
		}
	}
	if len(remember) == 0 {
		return
	}
	result := make([]int, 0, len(remember))
	for num := range remember {
		result = append(result, num)
	}
	sort.Ints(result)
	for i, num := range result {
		if i > 0 {
			writer.WriteString(" ")
		}
		writer.WriteString(strconv.Itoa(num))
	}
}
