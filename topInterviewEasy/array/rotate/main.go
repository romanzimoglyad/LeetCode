package main

import (
	"fmt"
)

func main() {
	in := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	rotate(in)
	fmt.Println(in)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	m := make(map[int][]int)
	for i := range matrix {
		for j := n - 1; j >= 0; j-- {
			m[i] = append(m[i], matrix[j][i])
		}
	}

	for i := range matrix {
		matrix[i] = m[i]
	}

}
