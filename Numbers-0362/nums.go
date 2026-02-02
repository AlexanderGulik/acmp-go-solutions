package main

import "fmt"

func main() {
	var n int
	result := 0
	fmt.Scan(&n)

	result = n - n/2 - n/3 - n/5 + n/6 + n/10 + n/15 - n/30

	fmt.Println(result)

}
