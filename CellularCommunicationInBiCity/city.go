package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	name := make([]string, 100)
	kvo := make([]int, 100)
	uk := -1
	x := make([]int, 10000)
	y := make([]int, 10000)
	r := make([]int, 10000)
	nk := make([]int, 10000)

	for z := 0; z < n; z++ {
		var a string
		fmt.Scan(&a)
		uk1 := -1
		for i := 0; i <= uk; i++ {
			if a == name[i] {
				uk1 = i
				break
			}
		}
		if uk1 == -1 {
			uk++
			name[uk] = a
			uk1 = uk
		}
		nk[z] = uk1
		fmt.Scan(&x[z], &y[z], &r[z])
	}
	var x0, y0 int
	fmt.Scan(&x0, &y0)

	for z := 0; z < n; z++ {
		if (x[z]-x0)*(x[z]-x0)+(y[z]-y0)*(y[z]-y0) <= r[z]*r[z] {
			kvo[nk[z]]++
		}
	}

	fmt.Println(uk + 1)
	for z := 0; z <= uk; z++ {
		fmt.Println(name[z], kvo[z])
	}
}
