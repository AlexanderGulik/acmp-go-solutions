package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	firstOne := n & -n
	nextZero := n + firstOne
	remOnes := n ^ nextZero
	remOnes = (remOnes / firstOne)
	remOnes >>= 2
	fmt.Println(nextZero | remOnes)
}

/*
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	target := count(n)
	for i := n + 1; ; i++ {
		if count(i) == target {
			fmt.Println(i)
			return
		}
	}
}

func count(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
*/

/*

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
		1. Переводим в двоичное разложение число n
		2. ПОдбераем число x > n, где x.count("1") == n.count("1") 
	countOnes := 0
	temp := n
	for temp > 0 {
		if temp&1 == 1 {
			countOnes++
		}
		temp >>= 1
	}
	x := n + 1
	for {
		countOnesX := 0
		tempX := x
		for tempX > 0 {
			if tempX&1 == 1 {
				countOnesX++
			}
			tempX >>= 1
		}
		if countOnesX == countOnes {
			fmt.Println(x)
			return
		} else {
			x++
		}
	}
}
*/
