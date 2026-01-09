package main

import (
	"fmt"
	"os"
)

func main() {
	in, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create("output.txt")
	if err != nil {
		return
	}
	defer out.Close()

	var n int
	fmt.Fscan(in, &n)

	var maxn int
	var t int
	var sum int64 = 0
	steak := []int{}
	var lastMin int

	fmt.Fscan(in, &t)
	steak = append(steak, t)
	lastMin = t
	maxn = t

	for z := 0; z < n-1; z++ {
		fmt.Fscan(in, &t)

		if t > lastMin {
			if t > maxn {
				maxn = t
			}
			for len(steak) > 1 {
				lastMin = steak[len(steak)-1]
				sum += int64(steak[len(steak)-2] - lastMin)
				lastMin = steak[len(steak)-2]
				steak = steak[:len(steak)-1]
			}
			sum += int64(t - lastMin)
			lastMin = maxn
			if len(steak) > 0 {
				steak = steak[:len(steak)-1]
			}
		}

		if t < lastMin {
			steak = append(steak, t)
			lastMin = t
		}
	}

	sum += int64(maxn - t)
	fmt.Fprint(out, sum)
}

/* я пытался решить это на гошке, но по времени никак не мог сделать лучше. 1.1c поэтому вот решение на ++. Я пытался его скопировать на го , но Time limit exceeded 3часа уже слишком
#include <fstream>

#include <vector>

using namespace std;



ifstream in("input.txt");

ofstream out("output.txt");



int main() {

    int n;        // количество интервалов

    in >> n;

    int maxn; // максимум

    int t;

    long long sum = 0;  // количество нужных динамиков

    vector <int> steak; // вектор будет хранить значение минимумов

    int lastMin;        // последний минимум

    in >> t;                // читаем первое данное

    steak.push_back(t);     // и помещаем его в стек как минимум

    lastMin = t;

    maxn = t;               // и запоминаем как максимум и последний минимум

    for (int z = 0; z < n - 1; z++) {

        in >> t;

        if (t > lastMin) {  // мы нашли повышение

            if (t > maxn) maxn = t;

            while (steak.size() > 1) { // находим сумму до ПЕРВОГО запомненного минимума

                lastMin = steak[steak.size() - 1];

                sum += steak[steak.size() - 2] - lastMin;

                lastMin = steak[steak.size() - 2]; // у нас новый последний минимум

                steak.pop_back();                  // удаляем отработанное значение

            }

            sum += t - lastMin;       // дополняем диапазон до запомненного минимального

            lastMin = maxn;           // и снова новый минимум

            if (steak.size() > 0) {

                steak.pop_back();             // очищаем стек

            }

        }

        if (t < lastMin) {  // новый минимум

            steak.push_back(t); // сохраняем данные в стеке

            lastMin = t;

        }

    }

    sum += maxn - t; // находим сумму до конца

    out << sum;

    return 0;

}
*/

// найти макс число
// добавлять числа с права и слева пока arr[n] == maxNun
// добавлять в отрезки по 1, пока arr[n] == maxNun
/*не хватает времени
var indxStart, indxEnd int
	for isTrueArr(arr, maxNun) {
		for indxStart < n && arr[indxStart] == maxNun {
			indxStart++
		}
		if indxStart >= n {
			break
		}

		indxEnd = indxStart
		for indxEnd < n && arr[indxEnd] != maxNun {
			indxEnd++
		}
		indxEnd--
		for i := indxStart; i <= indxEnd; i++ {
			arr[i]++
		}
		result++

	}

func isTrueArr(arr []int, maxNun int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] != maxNun {
			return true
		}
	}
	return false
}

тоже не рабочий вариант
	result := 0
	prev := 0
	for i := 0; i < n; i++ {
		if arr[i] > prev {
			result += arr[i] - prev
		}
		prev = arr[i]
	}
*/

/* ебал реп

import (
 "fmt"
 "os"
)

func main() {
 in, err := os.Open("input.txt")
 if err != nil {
  return
 }
 defer in.Close()

 out, err := os.Create("output.txt")
 if err != nil {
  return
 }
 defer out.Close()

 var n int
 fmt.Fscan(in, &n)

 var maxn int
 var t int
 var sum int64 = 0
 steak := []int{}
 var lastMin int

 fmt.Fscan(in, &t)
 steak = append(steak, t)
 lastMin = t
 maxn = t

 for z := 0; z < n-1; z++ {
  fmt.Fscan(in, &t)

  if t > lastMin {
   if t > maxn {
    maxn = t
   }
   for len(steak) > 1 {
    lastMin = steak[len(steak)-1]
    sum += int64(steak[len(steak)-2] - lastMin)
    lastMin = steak[len(steak)-2]
    steak = steak[:len(steak)-1]
   }
   sum += int64(t - lastMin)
   lastMin = maxn
   if len(steak) > 0 {
    steak = steak[:len(steak)-1]
   }
  }

  if t < lastMin {
   steak = append(steak, t)
   lastMin = t
  }
 }

 sum += int64(maxn - t)
 fmt.Fprint(out, sum)
}
*/
