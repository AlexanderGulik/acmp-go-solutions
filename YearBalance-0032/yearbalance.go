package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string
	var result int
	fmt.Scan(&a, &b)
	if b[0] == '-' {
		x := maxNumber(a)
		y := maxNumber(b[1:])
		result = x + y
	} else {
		x := maxNumber(a)
		y := minNumber(b)
		result = x - y
	}
	fmt.Println(result)
}

func minNumber(s string) int {
	b := []byte(s)
	for i := 0; i < len(b)-1; i++ {
		for j := i; j < len(b)-i-1; j++ {
			if !(i == 0 && b[j+1] == '0') {
				if b[i] < b[j+1] {
					b[i], b[j+1] = b[j+1], b[i]
				}
			}
		}
	}
	x, _ := strconv.Atoi(string(b))
	return x
}

func maxNumber(s string) int {
	b := []byte(s)
	for i := 0; i < len(b)-1; i++ {
		for j := i; j < len(b)-i-1; j++ {
			if !(i == 0 && b[j+1] == '0') {
				if b[i] < b[j+1] {
					b[i], b[j+1] = b[j+1], b[i]
				}
			}
		}
	}
	x, _ := strconv.Atoi(string(b))
	return x
}
