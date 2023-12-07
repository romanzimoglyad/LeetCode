package main

import (
	"fmt"
)

func main() {
	resp := comp2([][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}})
	fmt.Println(resp)

}

func comp(in [][]int) [][]int {
	n := len(in)
	matrix := make([][]int, n)

	for i, _ := range matrix {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][n-i-1] = in[i][j]
		}
	}
	return matrix
}

func comp2(in [][]int) [][]int {
	n := len(in)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			tmp := in[i][j]
			in[i][j] = in[n-j-1][i]
			in[n-j-1][i] = in[n-i-1][n-j-1]
			in[n-i-1][n-j-1] = in[j][n-i-1]
			in[j][n-i-1] = tmp
		}
	}
	return in
}
