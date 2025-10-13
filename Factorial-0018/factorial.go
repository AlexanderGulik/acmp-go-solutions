package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scan(&n)
	result := big.NewInt(1)
	for i := 1; i <= n; i++ {
		q := big.NewInt(int64(i))
		result.Mul(result, q)
	}
	fmt.Println(result)
}
