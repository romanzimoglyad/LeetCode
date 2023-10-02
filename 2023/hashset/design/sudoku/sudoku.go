package main

import "fmt"

func main() {

	in1 := [][]byte{
		{'.', '.', '.', '.', '.', '.', '5', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'9', '3', '.', '.', '2', '.', '4', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '3', '4', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '5', '2', '.', '.'}}

	t := isValidSudoku(in1)
	fmt.Println(t)
}

func isValidSudoku(board [][]byte) bool {
	for j := 0; j < len(board); j++ {
		mColumn := make(map[byte]struct{})
		for i := 0; i < len(board); i++ {
			if _, ok := mColumn[board[i][j]]; ok && board[i][j] != '.' {
				return false
			}
			mColumn[board[i][j]] = struct{}{}
		}
	}

	for j := 0; j < len(board); j++ {
		mRow := make(map[byte]struct{})
		for i := 0; i < len(board); i++ {
			if _, ok := mRow[board[j][i]]; ok && board[j][i] != '.' {
				return false
			}
			mRow[board[j][i]] = struct{}{}
		}
	}

	var rowEnd = 3
	var rowStart = 0
	for rowEnd <= 9 {
		var colStart = 0
		var colEnd = 3
		for colEnd <= 9 {
			mTriple := make(map[string]struct{})
			for i := rowStart; i < rowEnd; i++ {
				for j := colStart; j < colEnd; j++ {
					t := board[i][j]
					fmt.Sprintf("%c", t)
					if _, ok := mTriple[fmt.Sprintf("%c", t)]; ok && board[i][j] != '.' {
						return false
					}
					mTriple[fmt.Sprintf("%c", t)] = struct{}{}
				}
			}
			colStart += 3
			colEnd += 3
		}

		rowEnd += 3
		rowStart += 3
	}

	return true
}
