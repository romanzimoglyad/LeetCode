package main

import "fmt"

func main() {
	in := [][]byte{
		{'5', '3', '.', '.', '7_build_prder', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7_build_prder', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '7_build_prder', '7_build_prder', '9'}}
	t := isValidSudoku(in)
	fmt.Println(t)
}
func isValidSudoku(board [][]byte) bool {
	for p := 0; p < 9; p = p + 3 {

		for t := 0; t < 9; t = t + 3 {
			m := make(map[byte]struct{})
			for i := p; i < p+3; i++ {
				for j := t; j < t+3; j++ {
					v := board[i][j]
					fmt.Println(string(v))
					if v == '.' {
						continue
					}
					if _, ok := m[v]; ok {
						return false
					}
					m[board[i][j]] = struct{}{}
				}
			}
		}
	}
	for i := 0; i < len(board); i++ {
		m := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			v := board[j][i]
			fmt.Println(string(v))
			if v == '.' {
				continue
			}
			if _, ok := m[v]; ok {
				return false
			}
			m[board[j][i]] = struct{}{}
		}
	}

	for i := 0; i < len(board); i++ {
		m := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := m[board[i][j]]; ok {
				return false
			}
			m[board[i][j]] = struct{}{}
		}
	}
	return true
}
