/*
#include <bits/stdc++.h>

using namespace std;

void update_max(int& max1, int& max2, int val) {
    if (val > max1) {
        max2 = max1;
        max1 = val;
    } else if (val > max2) {
        max2 = val;
    }
}

void compute_for_row(int n, int row, int upper_idx, int lower_idx, int self_idx, vector<vector<int>>& lines, int& result) {
    for (int j = 0; j < n; j++) {
        int b = lines[self_idx][j];
        int max1 = numeric_limits<int>::min();
        int max2 = numeric_limits<int>::min();
        int count = 0;
        if (j > 0) {
            int val = lines[self_idx][j - 1];
            update_max(max1, max2, val);
            count++;
        }
        if (j < n - 1) {
            int val = lines[self_idx][j + 1];
            update_max(max1, max2, val);
            count++;
        }
        if (upper_idx >= 0) {
            int val = lines[upper_idx][j];
            update_max(max1, max2, val);
            count++;
        }
        if (lower_idx >= 0) {
            int val = lines[lower_idx][j];
            update_max(max1, max2, val);
            count++;
        }
        if (count < 2) continue;
        result = std::max(result, b + max1 + max2);
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    int n;
    cin >> n;
    if (n == 0) {
        cout << 0 << endl;
        return 0;
    }
    vector<vector<int>> lines(3, vector<int>(n));
    int result = numeric_limits<int>::min();
    int current_buffer = 0;
    for (int j = 0; j < n; j++) {
        cin >> lines[current_buffer][j];
    }
    if (n == 1) {
        compute_for_row(n, 0, -1, -1, current_buffer, lines, result);
        cout << result << endl;
        return 0;
    }
    int next_buffer = (current_buffer + 1) % 3;
    for (int j = 0; j < n; j++) {
        cin >> lines[next_buffer][j];
    }
    compute_for_row(n, 0, -1, next_buffer, current_buffer, lines, result);
    for (int row = 2; row < n; row++) {
        int prev_buffer = current_buffer;
        current_buffer = next_buffer;
        next_buffer = (current_buffer + 1) % 3;
        for (int j = 0; j < n; j++) {
            cin >> lines[next_buffer][j];
        }
        compute_for_row(n, row - 1, prev_buffer, next_buffer, current_buffer, lines, result);
    }
    compute_for_row(n, n - 1, current_buffer, -1, next_buffer, lines, result);
    cout << result << endl;
    return 0;
}
*/package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	tables := make([]int32, n*n)
	for i := 0; i < n*n; i++ {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		tables[i] = int32(val)
	}
	result := int32(math.MinInt32)
	dirs := [4][2]int8{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			b := tables[i*n+j]
			max1, max2 := int32(math.MinInt32), int32(math.MinInt32)
			count := 0
			for _, d := range dirs {
				ai, aj := i+int(d[0]), j+int(d[1])
				if ai < 0 || ai >= n || aj < 0 || aj >= n {
					continue
				}
				val := tables[ai*n+aj]
				count++
				if val > max1 {
					max2 = max1
					max1 = val
				} else if val > max2 {
					max2 = val
				}
			}
			if count < 2 {
				continue
			}
			result = max(result, b+max1+max2)
		}
	}
	fmt.Println(result)
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
