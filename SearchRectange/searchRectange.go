package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)

	left := make([]int, k+1)
	right := make([]int, k+1)
	up := make([]int, k+1)
	down := make([]int, k+1)

	for i := range left {
		left[i] = m + 1
		right[i] = -1
		up[i] = n + 1
		down[i] = -1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var a int
			fmt.Fscan(in, &a)

			if a == 0 {
				continue
			}

			y := n - 1 - i

			if j < left[a] {
				left[a] = j
			}
			if j > right[a] {
				right[a] = j
			}
			if y < up[a] {
				up[a] = y
			}
			if y > down[a] {
				down[a] = y
			}

			if j < left[0] {
				left[0] = j
			}
			if j > right[0] {
				right[0] = j
			}
			if y < up[0] {
				up[0] = y
			}
			if y > down[0] {
				down[0] = y
			}
		}
	}

	for i := 1; i <= k; i++ {
		if left[i] < m+1 {
			fmt.Fprintf(out, "%d %d %d %d\n",
				left[i], up[i], right[i]+1, down[i]+1)
		} else {
			fmt.Fprintf(out, "%d %d %d %d\n",
				left[0], up[0], right[0]+1, down[0]+1)
		}
	}
}

/*package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	field := make([][]int, n)
	for i := 0; i < n; i++ {
		field[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &field[i][j])
		}
	}
	results := make([][4]int, k+1)
	for p := 1; p <= k; p++ {
		results[p] = [4]int{m, n, -1, -1}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			p := field[i][j]
			if p > 0 && p <= k {
				y := n - 1 - i
				if j < results[p][0] {
					results[p][0] = j
				}

				if y < results[p][1] {
					results[p][1] = y
				}

				if j > results[p][2] {
					results[p][2] = j
				}

				if y > results[p][3] {
					results[p][3] = y
				}
			}
		}
	}

	for i := 1; i <= k; i++ {
		if results[i][2] == -1 {
			fmt.Fprintf(out, "0 0 0 0\n")
		} else {
			fmt.Fprintf(out, "%d %d %d %d\n", results[i][0], results[i][1], results[i][2]+1, results[i][3]+1)
		}
	}

}*/
