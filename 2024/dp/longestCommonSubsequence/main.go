package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("abcdea", "acya")) // aca
}

// a
// a cya

// bcdea
// c ya

// b cdea
// cya

func longestCommonSubsequence(text1 string, text2 string) int {
	m = make(map[[2]int]int)
	return dp(0, 0, text1, text2)
}

var m = map[[2]int]int{}

func dp(i, j int, text1 string, text2 string) int {
	if i >= len(text1) || j >= len(text2) {
		return 0
	}
	if _, ok := m[[2]int{i, j}]; !ok {
		if text1[i] == text2[j] {
			m[[2]int{i, j}] = dp(i+1, j+1, text1, text2) + 1
		} else {
			m[[2]int{i, j}] = max(dp(i+1, j, text1, text2), dp(i, j+1, text1, text2))
		}
	}

	return m[[2]int{i, j}]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
