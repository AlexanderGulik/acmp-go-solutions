package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var rankMapping = map[byte]int{
	'6': 0, '7': 1, '8': 2, '9': 3,
	'T': 4, 'J': 5, 'Q': 6, 'K': 7, 'A': 8,
}

var suitMapping = map[byte]int{
	'S': 0, 'C': 1, 'D': 2, 'H': 3,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := strings.Fields(scanner.Text())
	trump := input[2][0]

	playerCards := make([][]int, 4)
	for i := range playerCards {
		playerCards[i] = make([]int, 9)
	}

	attackCards := make([][]int, 4)
	for i := range attackCards {
		attackCards[i] = make([]int, 9)
	}

	scanner.Scan()
	playerHand := strings.Fields(scanner.Text())
	for _, card := range playerHand {
		rank := rankMapping[card[0]]
		suit := suitMapping[card[1]]
		playerCards[suit][rank] = 1
	}

	scanner.Scan()
	attackHand := strings.Fields(scanner.Text())
	for _, card := range attackHand {
		rank := rankMapping[card[0]]
		suit := suitMapping[card[1]]
		attackCards[suit][rank] = 1
	}

	trumpSuit := suitMapping[trump]

	possibleToDefend := true

	for suit := 0; suit < 4; suit++ {
		for rank := 0; rank < 9; rank++ {
			if attackCards[suit][rank] == 1 {
				found := false
				for nextRank := rank + 1; nextRank < 9; nextRank++ {
					if playerCards[suit][nextRank] == 1 {
						playerCards[suit][nextRank] = 0
						found = true
						break
					}
				}

				if !found && suit != trumpSuit {
					for kozyrRank := 0; kozyrRank < 9; kozyrRank++ {
						if playerCards[trumpSuit][kozyrRank] == 1 {
							playerCards[trumpSuit][kozyrRank] = 0
							found = true
							break
						}
					}
				}

				if !found {
					possibleToDefend = false
					break
				}
			}
		}

		if !possibleToDefend {
			break
		}
	}

	if possibleToDefend {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

/*package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	var char byte

	fmt.Scan(&n, &m)
	fmt.Scan(&char)
	player := []string{}
	attack := []string{}
	for i := 0; i < n; i++ {
		var curr string
		fmt.Scan(&curr)
		player = append(player, curr)
	}
	for i := 0; i < m; i++ {
		var curr string
		fmt.Scan(&curr)
		attack = append(attack, curr)
	}
	sortCards(player)
	sortCards(attack)
	for i := 0; i < len(attack); i++ {
		for j := 0; j < len(player); j++ {
			if (attack[i][1] == player[j][1] && attack[i][0] < player[j][0]) ||
				(byte(player[j][1]) == char && byte(attack[i][1]) != char) {
				attack = removeElement(attack, i)
				player = removeElement(player, j)
				i--
				j--
				break

			}
		}
	}

	if len(attack) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

func sortCards(cards []string) {
	sort.Slice(cards, func(i, j int) bool {

		ranks := map[rune]int{
			'6': 0, '7': 1, '8': 2,
			'9': 3, 'T': 4, 'J': 5,
			'Q': 6, 'K': 7, 'A': 8}

		ri, rj := ranks[rune(cards[i][0])], ranks[rune(cards[j][0])]
		si, sj := cards[i][1], cards[j][1]

		if ri != rj {
			return ri < rj
		}
		return si < sj
	})
}

func removeElement(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}*/
