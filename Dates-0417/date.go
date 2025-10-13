package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Scan(&n)
	startDate := time.Date(2008, time.January, 1, 0, 0, 0, 0, time.UTC)
	future := startDate.AddDate(0, 0, n)
	fmt.Printf("%s, %s", future.Weekday(), future.Format("02.01"))
}
