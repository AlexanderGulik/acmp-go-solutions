package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	move := -1
	for i := 0; i < k; i++ {
		zero, one := 0, 0
		pos0, pos1 := -1, -1
		for j := i; j < n; j += k {
			if arr[j] == 0 {
				zero++
				if zero == 1 {
					pos0 = j
				}
			} else {
				one++
				if one == 1 {
					pos1 = j
				}
			}
		}
		if zero == 0 || one == 0 {
			continue
		}

		if (zero > 1 && one > 1) || move != -1 {
			fmt.Println("FAIL")
			return
		}
		if zero == 1 && one == 1 {
			if pos0 < pos1 {
				move = pos0
			} else {
				move = pos1
			}
		} else if zero == 1 {
			move = pos0
		} else if one == 1 {
			move = pos1
		}
	}

	fmt.Println("OK")
	fmt.Println(move + 1)

}
