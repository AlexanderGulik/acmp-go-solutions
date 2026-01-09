package main

import "fmt"

func main() {
	var s, k int
	fmt.Scan(&s, &k)
	num1 := make([]int, k)
	num2 := make([]int, k)
	// Первые два числа первого спутника представляют собой сумму цифр и количество цифр тех двух чисел, которыми должен ответить второй спутник.
	// При этом в качестве ответа должны получиться числа, представляющие наибольшее и наименьшее возможные значения, которые могут быть сформированы по описанному выше методу.
	// что можем заметить число k == кол-во цифр в числах s k
	// число s = начало числа num1 || кол-во нулей
	// первое число num1 = s && окончание num2 = s-1
	// ошибка s != цифра 
	tempMaxS := s
	indMax := 0
	for tempMaxS != 0 {
		if num1[indMax] != 9 {
			num1[indMax]++
			tempMaxS--
		} else {
			indMax++
		}
	}

	indMin := k - 1
	tempMinS := s - 1
	num2[0] = 1
	for tempMinS != 0 {
		if num2[indMin] != 9 {
			num2[indMin]++
			tempMinS--
		} else {
			indMin--
		}
	}

	for i := 0; i < k; i++ {
		fmt.Printf("%d", num1[i])
	}
	fmt.Printf(" ")
	for i := 0; i < k; i++ {
		fmt.Printf("%d", num2[i])
	}

}
