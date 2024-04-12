package main

import "fmt"

func main() {
	arr := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	fmt.Println(maximalSquare(arr))

	arr1 := [][]byte{
		{'1', '1'},
		{'1', '1'},
	}
	fmt.Println(maximalSquare(arr1))
}

func maximalSquare(matrix [][]byte) int {

	arr := make([][]int, len(matrix)+1)
	for i := range arr {
		arr[i] = make([]int, len(matrix[0])+1)
	}

	max := 0
	for i := 1; i < len(matrix)+1; i++ {
		for j := 1; j < len(matrix[0])+1; j++ {
			if matrix[i-1][j-1] == '1' {
				arr[i][j] = minimum(minimum(arr[i-1][j], arr[i][j-1]), arr[i-1][j-1]) + 1
				if arr[i][j] > max {
					max = arr[i][j]
				}
			}
		}
	}

	return max * max
}

func minimum(i, j int) int {
	if i < j {
		return i
	}
	return j
}
