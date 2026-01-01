package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	n--
	a := "0"
	for i := 0; i < n; i++ {
		a = f(a)
	}
	fmt.Println(a)
}

func f(a string) string {
	v := make([]int, 10)
	for i := 0; i < len(a); i++ {
		v[a[i]-'0']++
	}
	p := -1
	for i := 0; i < len(v); i++ {
		if v[i] == 0 {
			p = i
			break
		}
	}
	f := -1
	start := int(a[0] - '0')
	if start < 1 {
		start = 1
	}
	for i := start; i < 10; i++ {
		if v[i] == 0 {
			f = i
			break
		}
	}
	b := []byte(a)
	for i := 0; i < len(a)-1; i++ {
		b[i+1] = byte(p) + '0'
	}

	if f == -1 {
		for i := 1; i < 10; i++ {
			if v[i] == 0 {
				f = i
				break
			}
		}
		b[0] = byte(f) + '0'
		b = append(b, byte(p)+'0')
	} else {
		b[0] = byte(f) + '0'
	}
	return string(b)
}

/*package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	num := make([]int, n+1)
	num[1] = 1
	for i := 1; i < n; i++ {
		num[i] = num[i-1]
		for !checkNumbers(num[i], num[i-1]) {
			num[i]++
		}
	}
	fmt.Println(num[n-1])
}

func checkNumbers(x, num int) bool {
	strA := strconv.Itoa(x)
	strB := strconv.Itoa(num)
	for _, charA := range strA {
		if strings.ContainsRune(strB, charA) {
			return false
		}
	}
	return true
}
*/
