package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	inputFile, _ := os.Open("INPUT.TXT")

	defer inputFile.Close()

	outputFile, _ := os.Create("OUTPUT.TXT")

	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var m, n int
	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			m, _ = strconv.Atoi(parts[0])
			n, _ = strconv.Atoi(parts[1])
		}
	}

	var k int
	if scanner.Scan() {
		k, _ = strconv.Atoi(scanner.Text())
	}

	maps := make(map[int][]int)

	for i := 0; i < k && scanner.Scan(); i++ {
		line := scanner.Text()
		lineSplit := strings.Fields(line)
		if len(lineSplit) == 0 {
			continue
		}
		curr := lineSplit[0]
		switch curr {
		case "ADD":
			if len(lineSplit) >= 3 {
				element, err := strconv.Atoi(lineSplit[1])
				if err != nil {
					continue
				}
				set, err := strconv.Atoi(lineSplit[2])
				if err != nil {
					continue
				}
				if element < 1 || element > m || set < 1 || set > n {
					continue
				}

				if !searchElement(maps[set], element) {
					maps[set] = append(maps[set], element)
				}
			}
		case "LISTSET":
			if len(lineSplit) >= 2 {
				set, err := strconv.Atoi(lineSplit[1])
				if err != nil || set < 1 || set > n {
					fmt.Fprintf(outputFile, "-1\n")
					continue
				}
				if len(maps[set]) == 0 {
					fmt.Fprintf(outputFile, "-1\n")
				} else {
					elements := make([]int, len(maps[set]))
					copy(elements, maps[set])
					sort.Ints(elements)
					for i, data := range elements {
						if i > 0 {
							fmt.Fprintf(outputFile, " ")
						}
						fmt.Fprintf(outputFile, "%d", data)
					}
					fmt.Fprintf(outputFile, "\n")
				}
			}
		case "LISTSETSOF":
			if len(lineSplit) >= 2 {
				element, err := strconv.Atoi(lineSplit[1])
				if err != nil || element < 1 || element > m {
					fmt.Fprintf(outputFile, "-1\n")
					continue
				}
				sets := []int{}
				for set, data := range maps {
					if searchElement(data, element) {
						sets = append(sets, set)
					}
				}
				if len(sets) == 0 {
					fmt.Fprintf(outputFile, "-1\n")
				} else {
					sort.Ints(sets)
					for i, set := range sets {
						if i > 0 {
							fmt.Fprintf(outputFile, " ")
						}
						fmt.Fprintf(outputFile, "%d", set)
					}
					fmt.Fprintf(outputFile, "\n")
				}
			}
		}
	}
}

func searchElement(arr []int, n int) bool {
	for _, data := range arr {
		if data == n {
			return true
		}
	}
	return false
}
