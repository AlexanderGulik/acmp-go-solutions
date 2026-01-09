package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	result := make([]int, n)
	flow := 0
	for i := 0; i < n; i++ {
		//авто на право = не едеет через мост. авто на лево - всегда едет через мост. авто прямо - зависит от направления
		//необходимо посчитать поток авто, которые поедут через мосты объездной дороги
		// считаем только мосты по объездной дороги
		// L - означает что об дорога идет по земле
		// B - означает что об дорога идет по мосту +
		// первые 2 числа - кол-во авто покинувших об дорогу, первое на лево, второе на право
		// вторые 2 числа - кол-во авто выехавших на об дорогу первое на лево, второе направо.

		var char string
		fmt.Scan(&char)
		var entLeft, entRight, exitLeft, exitRight int
		fmt.Scan(&exitLeft, &exitRight, &entLeft, &entRight)
		if char == "B" {
			current := flow - exitRight + entLeft
			result[i] = current
		} else if char == "L" {
			result[i] = -1
		}
		flow += entLeft + entRight - exitLeft - exitRight
	}

	for i := 0; i < len(result); i++ {

		fmt.Printf("%d ", result[i])

	}

}

// задача была простой, я просто еблан и попутал ent and exit в вводе данных и из-за этого страдал хуйней
