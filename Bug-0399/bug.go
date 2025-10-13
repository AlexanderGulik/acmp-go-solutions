package main

import (
	"bufio"
	"fmt"
	"os"
)

type Napr int

const (
	downn Napr = iota
	rightn
	upn
	leftn
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	if n < 1 || m < 1 || n > 100 || m > 100 {
		fmt.Println(-1)
		return
	}

	ss := make([][]int, n+1)
	for i := range ss {
		ss[i] = make([]int, m+1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	s := make([]string, n+1)
	for z := 1; z <= n; z++ {
		if !scanner.Scan() {
			fmt.Println(-1)
			return
		}
		line := scanner.Text()
		if len(line) != m {
			fmt.Println(-1)
			return
		}
		s[z] = "@" + line
	}

	bugn := downn
	steps := 0
	x, y := 2, 2
	olddx, olddy := 0, 1
	ss[2][2] = 1

	for x != m-1 || y != n-1 {
		var dx, dy int
		var kuda Napr
		z := 2000000

		if y+1 <= n && s[y+1][x] != '@' && ss[y+1][x] < z {
			z = ss[y+1][x]
			kuda = downn
			dx, dy = 0, 1
		}
		if x+1 <= m && s[y][x+1] != '@' && ss[y][x+1] < z {
			z = ss[y][x+1]
			kuda = rightn
			dx, dy = 1, 0
		}
		if y-1 >= 1 && s[y-1][x] != '@' && ss[y-1][x] < z {
			z = ss[y-1][x]
			kuda = upn
			dx, dy = 0, -1
		}
		if x-1 >= 1 && s[y][x-1] != '@' && ss[y][x-1] < z {
			z = ss[y][x-1]
			kuda = leftn
			dx, dy = -1, 0
		}

		if steps > 10000000 || z == 2000000 {
			fmt.Println(-1)
			return
		}

		if kuda == bugn {
			x += dx
			y += dy
			steps++
			ss[y][x]++
		} else {
			nx, ny := x+olddx, y+olddy
			if ny >= 1 && ny <= n && nx >= 1 && nx <= m && s[ny][nx] != '@' && ss[y+dy][x+dx] == ss[ny][nx] {
				x = nx
				y = ny
				steps++
				kuda = bugn
				ss[y][x]++
			} else {
				bugn = kuda
				x += dx
				y += dy
				steps++
				ss[y][x]++
				olddx, olddy = dx, dy
			}
		}
	}

	fmt.Println(steps)
}

/*package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	in, _ := os.Open("input.txt")
	out, _ := os.Create("output.txt")
	defer in.Close()
	defer out.Close()

	reader := bufio.NewReader(in)
	writer := bufio.NewWriter(out)
	defer writer.Flush()
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	maze := make([]string, n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		maze[i] = line[:len(line)-1]
	}

	if maze[n-2][m-2] == '@' {
		fmt.Fprintln(writer, -1)
		return
	}
	visits := make([][]int, n)
	for i := range visits {
		visits[i] = make([]int, m)
	}

	directions := []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	x, y := 1, 1
	visits[x][y] = 1
	steps := 0
	dir := 0
	for x != n-2 || y != m-2 {
		if steps > 10000000 {
			fmt.Fprintln(writer, -1)
			return
		}
		minVisits := -1
		candidateDirs := []int{}
		for d := 0; d < 4; d++ {
			nx, ny := x+directions[d].x, y+directions[d].y
			if maze[nx][ny] != '@' {
				if minVisits == -1 || visits[nx][ny] < minVisits {
					minVisits = visits[nx][ny]
					candidateDirs = []int{d}
				} else if visits[nx][ny] == minVisits {
					candidateDirs = append(candidateDirs, d)
				}
			}
		}
		newDir := -1

		for _, d := range candidateDirs {
			if d == dir {
				newDir = d
				break
			}
		}
		if newDir == -1 {
			for _, priorityDir := range []int{0, 1, 2, 3} {
				for _, d := range candidateDirs {
					if d == priorityDir {
						newDir = d
						break
					}
				}
				if newDir != -1 {
					break
				}
			}
		}
		dir = newDir
		x += directions[dir].x
		y += directions[dir].y
		visits[x][y]++
		steps++
	}
	fmt.Fprintln(writer, steps)
}*/
