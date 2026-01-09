package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	b := big.NewInt(0)
	var ia, ib string
	fmt.Scan(&ia, &ib)
	a.SetString(ia, 10)
	b.SetString(ib, 10)
	result := new(big.Int).Mul(a, b)
	fmt.Println(result)
}
