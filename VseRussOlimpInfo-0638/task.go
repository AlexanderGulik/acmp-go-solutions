package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var w, dw, s int
	fmt.Scan(&w, &dw, &s)
	weekForbidden := make([]bool, w+1)

	for i := 0; i < dw; i++ {
		var day int
		fmt.Scan(&day)
		weekForbidden[day] = true
	}
	var dm int
	fmt.Scan(&dm)

	monthForbidden := make([]bool, n+1)
	for i := 0; i < dm; i++ {
		var day int
		fmt.Scan(&day)
		monthForbidden[day] = true
	}

	allowed := make([]int, n+1)
	for i := 1; i <= n; i++ {
		weekDay := ((i-1)+(s-1))%w + 1
		if !weekForbidden[weekDay] && !monthForbidden[i] {
			allowed[i] = 1
		}
	}

	pref := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pref[i] = pref[i-1] + allowed[i]
	}
	result := 0

	for i := 1; i <= n-k+1; i++ {
		if pref[i+k-1]-pref[i-1] == k {
			result++
		}
	}
	fmt.Println(result)
}
