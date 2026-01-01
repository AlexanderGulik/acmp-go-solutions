package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	set := make(map[int]bool)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		set[num] = true
	}

	a1 := make([]int, 0, len(set))
	for num := range set {
		a1 = append(a1, num)
	}
	sort.Ints(a1)

	n = len(a1)
	result := ""
	t := 0

	for t < n-2 && n > 2 {
		if a1[t]+1 == a1[t+1] && a1[t+1]+1 == a1[t+2] {
			k := t + 1
			for k <= n-2 && a1[k]+1 == a1[k+1] {
				k++
			}

			s1 := fmt.Sprintf("%d, ..., %d", a1[t], a1[k])
			s2 := ""
			for z := t; z <= k; z++ {
				s2 += strconv.Itoa(a1[z]) + ", "
			}
			s2 = s2[:len(s2)-2]

			if len(s1) <= len(s2) {
				result += s1
			} else {
				result += s2
			}

			t = k + 1
			if t < n {
				result += ", "
			}
		} else {
			result += strconv.Itoa(a1[t]) + ", "
			t++
		}
	}

	if t < n-1 {
		result += strconv.Itoa(a1[t]) + ", "
		t++
	}
	if t == n-1 {
		result += strconv.Itoa(a1[t])
	}

	fmt.Print(result)
}
