/*
	Просто идите нахуй ебучий компилятор acmp алгоритм одинковый, но на ++ принимает, на гошке только хуй в рот

#include <bits/stdc++.h>

using namespace std;

	int main() {
	    int h;
	    cin >> h;
	    vector<int> line(h);
	    for (int i = 0; i < h; i++) {
	        cin >> line[i];
	    }
	    string str;
	    cin >> str;
	    int L = str.length();
	    int w = L / h;
	    int r = L % h;
	    int start = 0;
	    vector<vector<char>> matrix(h);
	    for (int i = 0; i < h; i++) {
	        int a = line[i] - 1;
	        int clen = w;
	        if (r != 0 && a < r) {
	            clen = w + 1;
	        }
	        string chunk = str.substr(start, clen);
	        start += clen;
	        matrix[a].resize(clen);
	        for (int j = 0; j < clen; j++) {
	            matrix[a][j] = chunk[j];
	        }
	    }
	    int max_col = 0;
	    for (auto& row : matrix) {
	        if (row.size() > max_col) {
	            max_col = row.size();
	        }
	    }
	    string result = "";
	    for (int col = 0; col < max_col; col++) {
	        for (int row = 0; row < h; row++) {
	            if (col < matrix[row].size()) {
	                result += matrix[row][col];
	            }
	        }
	    }
	    cout << result << endl;
	    return 0;
	}
*/
package main

import "fmt"

func main() {
	var h int
	fmt.Scan(&h)
	line := make([]int, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&line[i])
	}
	var str string
	fmt.Scan(&str)
	runes := []rune(str)
	L := len(runes)
	w := L / h
	r := L % h
	var start int = 0
	matrix := make([][]rune, h)
	for i := 0; i < h; i++ {
		a := line[i] - 1
		clen := w
		if a < r {
			clen = w + 1
		}
		chunk := runes[start : start+clen]
		start += clen
		matrix[a] = make([]rune, clen)
		copy(matrix[a], chunk)
	}
	max_col := 0
	for _, row := range matrix {
		if len(row) > max_col {
			max_col = len(row)
		}
	}
	var res []rune
	for col := 0; col < max_col; col++ {
		for row := 0; row < h; row++ {
			if col < len(matrix[row]) {
				res = append(res, matrix[row][col])
			}
		}
	}
	result := string(res)
	fmt.Println(result)
}

/*package main

import "fmt"

func main() {
	var h int
	fmt.Scan(&h)
	line := make([]int, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&line[i])
	}
	var str string
	fmt.Scan(&str)
	runes := []rune(str)
	lenStr := (len(runes) + h - 1) / h
	matrix := make([][]rune, h)
	for i := 0; i < h; i++ {
		matrix[i] = make([]rune, lenStr)
		for j := 0; j < lenStr; j++ {
			matrix[i][j] = ' '
		}
	}

	for i := 0; i < h; i++ {
		start := i * lenStr
		end := start + lenStr
		if end > len(runes) {
			end = len(runes)
		}
		chunk := runes[start:end]

		a := line[i] - 1
		if a >= 0 && a < h {
			for j := 0; j < len(chunk); j++ {
				if j < len(matrix[a]) {
					matrix[a][j] = rune(chunk[j])
				}
			}
		}
	}
	var result string
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			result += string(matrix[j][i])
		}
	}
	fmt.Println(result)
}*/
