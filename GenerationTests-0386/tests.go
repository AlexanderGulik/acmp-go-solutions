// даунская парабола тащит
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 1 || n > 300 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	for i := 0; i < n; i++ {
		x := i
		y := (i * i * 179) % 9973
		fmt.Println(x, y)
	}
}
