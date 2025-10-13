package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	table := make([][]rune, n)
	for i, _ := range table {
		table[i] = make([]rune, n)
	}

	for i := 0; i < n; i++ {
		var text string
		fmt.Scan(&text)
		for j := 0; j < n; j++ {
			table[i][j] = rune(text[j])
		}
	}
	for i := 0; i < m; i++ {
		var text string
		fmt.Scan(&text)
		for a := 0; a < len(text); a++ {
			searchTable(table, rune(text[a]))
		}
	}
	result := []rune{}
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[0]); j++ {
			if table[i][j] != '0' {
				result = append(result, table[i][j])
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	for i := 0; i < len(result); i++ {
		fmt.Printf(string(result[i]))
	}
}

func searchTable(table [][]rune, a rune) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			if table[i][j] == a {
				table[i][j] = '0'
			}
		}
	}
}
