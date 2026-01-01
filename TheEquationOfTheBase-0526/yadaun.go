package main

import "fmt"

func main() {
	var a int
	var b string
	fmt.Scan(&b, &a)
	s := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for z := 2; z < 37; z++ {
		a1 := a
		c := ""
		for a1 > 0 {
			c = string(s[a1%z]) + c
			a1 = a1 / z
		}
		if b == c {
			fmt.Print(z)
			return
		}
	}
	fmt.Print(0)
}
