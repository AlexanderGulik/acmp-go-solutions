package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	result := merge(s1, s2)
	fmt.Println(result)
}

func merge(a, b string) string {
	result := make([]byte, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] || (a[i] == b[j] && cmp(a[i:], b[j:]) < 0) {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	for i < len(a) {
		result = append(result, a[i])
		i++
	}
	for j < len(b) {
		result = append(result, b[j])
		j++
	}
	return string(result)
}

func cmp(a, b string) int {
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	if i == len(a) {
		return 1
	}
	if i == len(b) {
		return -1
	}
	if a[i] < b[i] {
		return -1
	}
	return 1
}

/*
func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	result := make([]byte, 0, len(s1)+len(s2))
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			result = append(result, s1[i])
			i++
		} else {
			result = append(result, s2[j])
			j++
		}
	}
	for i < len(s1) {
		result = append(result, s1[i])
		i++
	}

	for j < len(s2) {
		result = append(result, s2[j])
		j++
	}
	fmt.Println(string(result))
}
*/
/*
package main

import (
	"fmt"
	"sort"
)

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	str := s1 + s2
	result := []byte(str)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	fmt.Println(string(result))
}*/
