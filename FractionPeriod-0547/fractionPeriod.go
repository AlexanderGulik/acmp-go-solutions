package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	prepreiod := 0
	period := 0
	temp := n
	count2 := 0
	for temp%2 == 0 {
		temp /= 2
		count2++
	}
	count5 := 0
	for temp%5 == 0 {
		temp /= 5
		count5++
	}
	if count2 > count5 {
		prepreiod = count2
	} else {
		prepreiod = count5
	}

	if temp == 1 {
		if prepreiod > 0 {
			period = 1
		} else {
			period = 0
		}
	} else {
		remainder := 1
		seen := make([]int, temp)
		for i := 1; ; i++ {
			remainder = (remainder * 10) % temp
			if seen[remainder] != 0 {
				period = i - seen[remainder]
				break
			}
			seen[remainder] = i
			if remainder == 0 {
				period = 0
				break
			}
		}
	}
	fmt.Println(prepreiod, period)
}
