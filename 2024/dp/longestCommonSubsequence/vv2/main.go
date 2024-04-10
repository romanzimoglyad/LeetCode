package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("aabcaad", "abc"))
}

func longestCommonSubsequence(text1 string, text2 string) int {

	arr := make([][]int, len(text1)+1)

	for i := range arr {
		arr[i] = make([]int, len(text2)+1)
	}
	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				arr[i][j] = arr[i+1][j+1] + 1
			} else {
				arr[i][j] = max(arr[i+1][j], arr[i][j+1])
			}
		}
	}

	return arr[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
