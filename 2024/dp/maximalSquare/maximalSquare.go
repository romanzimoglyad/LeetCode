package main

import "fmt"

func main() {
	arr := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	fmt.Println(maximalSquare(arr))
}

func maximalSquare(matrix [][]byte) int {
	var res int
	rows := len(matrix)
	var cols int
	if rows != 0 {
		cols = len(matrix[0])
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			var tmpRes int

			if matrix[i][j] == '1' {
				tmpRes = 1
				newI := i + 1
				newJ := j + 1
				for newI < rows && newJ < cols {
					correct := true
					if matrix[newI][newJ] == '1' {
						for k := i; k < newI; k++ {
							if matrix[k][newJ] == '0' {
								correct = false
								break
							}
						}
						if correct {
							for k := j; k < newJ; k++ {
								if matrix[newI][k] == '0' {
									correct = false
									break
								}
							}
						}
					} else {
						break
					}
					if correct {
						newI = newI + 1
						newJ = newJ + 1
						tmpRes++
					} else {
						break
					}

				}
			}
			res = max(res, tmpRes*tmpRes)
		}
	}

	return res
}
