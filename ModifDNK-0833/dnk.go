package main

import "fmt"

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	runesS1 := []rune(s1)
	runesS2 := []rune(s2)
	count := 0
	result := make([][2]int, 0)
	for i := 0; i < len(runesS1); i++ {
		if runesS1[i] != runesS2[i] {
			for j := i; j < len(runesS2); j++ {
				if runesS1[j] == runesS2[i] {
					reverse(runesS1, i, j)
					result = append(result, [2]int{i + 1, j + 1})
					count++
					break
				}
			}
		}
	}
	fmt.Println(count)
	for i := 0; i < count; i++ {
		fmt.Println(result[i][0], result[i][1])
	}
}

func reverse(s []rune, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
