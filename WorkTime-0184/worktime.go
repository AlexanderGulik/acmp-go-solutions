package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	const (
		restTime = 16 * 60
		dayTime  = 24 * 60
	)

	d := make([]int, n)
	for z := 0; z < n; z++ {
		scanner.Scan()
		s := strings.TrimSpace(scanner.Text())
		if len(s) < 12 {
			return
		}
		day := int((s[0]-'0')*10 + (s[1] - '0'))
		month := int((s[3]-'0')*10 + (s[4] - '0'))
		hour := int((s[7]-'0')*10 + (s[8] - '0'))
		minute := int((s[10]-'0')*10 + (s[11] - '0'))
		t := hour*60 + minute
		d[z] = (jul(2017, month, day)-jul(2017, 1, 1))*dayTime + t
	}

	sort.Ints(d)

	ss := 0
	for z := 0; z < n; z += 2 {
		for d[z+1]-d[z] >= dayTime {
			d[z+1] -= dayTime
			ss += 480
		}
		if d[z+1]-d[z] >= restTime {
			d[z+1] -= restTime
		}
		ss += d[z+1] - d[z] + 1
	}

	fmt.Printf("%d:", ss/60)
	if ss%60 < 10 {
		fmt.Print("0")
	}
	fmt.Print(ss % 60)
}

func jul(year, month, day int) int {
	a := (14 - month) / 12
	y := year + 4800 - a
	m := month + 12*a - 3
	return day + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045
}
