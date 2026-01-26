package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	s   string
	pos int
	c   byte
)

const EOT = '\n'

func nextChar() {
	if pos < len(s) {
		c = s[pos]
		pos++
	} else {
		c = EOT
	}
}

func no() {
	fmt.Println("NO")
	os.Exit(0)
}

func number() {
	if c == '1' {
		nextChar()
		if '0' <= c && c <= '9' {
			nextChar()
		} else {
			no()
		}
		for '0' <= c && c <= '9' {
			nextChar()
		}
	} else if '2' <= c && c <= '9' {
		nextChar()
		for '0' <= c && c <= '9' {
			nextChar()
		}
	} else {
		no()
	}
}

func element() string {
	elm := []byte{}
	if 'A' <= c && c <= 'Z' {
		elm = append(elm, c)
		nextChar()
	} else {
		no()
	}
	if 'a' <= c && c <= 'z' {
		elm = append(elm, c)
		nextChar()
	}
	if '0' <= c && c <= '9' {
		number()
	}
	return string(elm)
}

func formula() {
	prev := element()
	for c != EOT {
		cur := element()
		if prev == cur {
			no()
		}
		prev = cur
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &s)
	pos = 0
	nextChar()
	formula()
	if c != EOT {
		no()
	}
	fmt.Println("YES")
}

/* я слаб и не смог решить данную задачу
package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	for i := 0; i < len(str)-1; i++ {
		if isDigit(rune(str[i])) && isDigit(rune(str[i+1])) {
			fmt.Println("NO")
			return
		}
	}
	for i := 0; i < len(str); i++ {
		if !isUpper(rune(str[i])) {
			fmt.Println("NO")
			return
		}

		elemStart := i
		i++
		if i < len(str) && isLower(rune(str[i])) {
			i++
		}

		for i < len(str) && isDigit(rune(str[i])) {
			i++
		}

		if i < len(str) && isUpper(rune(str[i])) {
			nextElemStart := i
			nextElemEnd := i + 1
			if nextElemEnd < len(str) && isLower(rune(str[nextElemEnd])) {
				nextElemEnd++
			}
			curElemEnd := elemStart + 1

			if curElemEnd < len(str) && isLower(rune(str[curElemEnd])) {
				curElemEnd++
			}

			curElem := str[elemStart:curElemEnd]
			nextElem := str[nextElemStart:nextElemEnd]
			if curElem == nextElem {
				fmt.Println("NO")
				return
			}

		}
		i--
	}
	fmt.Println("YES")
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
*/
