package main

import "fmt"

var may map[int]int

func f(n int) int {
	if n < 3 {
		return 0
	}
	if n == 3 {
		return 1
	}
	if val, ex := may[n]; ex {
		return val
	}
	even := n / 2
	odd := n - even
	result := f(even) + f(odd)
	may[n] = result
	return result
}
func main() {
	var n int
	fmt.Scan(&n)

	may = make(map[int]int)

	result := f(n)
	fmt.Println(result)
}
