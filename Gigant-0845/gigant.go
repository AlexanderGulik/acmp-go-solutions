package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x, y string
	fmt.Scan(&x, &y)
	x1 := int(x[len(x)-1] - '0')
	var yMod100 int
	if len(y) >= 2 {
		yMod100, _ = strconv.Atoi(y[len(y)-2:])
	} else {
		yMod100, _ = strconv.Atoi(y)
	}
	if y == "0" {
		fmt.Println(1)
		return
	}

	yMod100 += 100

	result := 1
	for i := 0; i < yMod100; i++ {
		result = (result * x1) % 10
	}
	fmt.Println(result)
}
