package main

import (
	"fmt"
)

func main() {
	var x, m, l, v int
	fmt.Scan(&x, &m, &l, &v)
	s := make([]byte, l)
	for i := 0; i < l; i++ {
		s[i] = '0'
	}
	start := string(s)
	current := start
	for calculateHash(current, x, m) != v {
		current = nextString(current)
		if current == string(s) {
			fmt.Println("NO SOLUTION")
			return
		}
	}
	fmt.Println(current)
}

func calculateHash(s string, x, m int) int {
	h := 0
	r := 1
	l := len(s)
	for i := 0; i < l; i++ {
		digit := int(s[i] - '0')
		h = (h + digit*r) % m
		r = (r * x) % m
	}
	return h
}

func nextString(s string) string {
	t := []byte(s)
	i := len(t) - 1
	for i >= 0 && t[i] == '9' {
		t[i] = '0'
		i--
	}
	if i >= 0 {
		t[i]++
	}
	return string(t)
}
