package main

import "fmt"

func main() {
	arr := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '1'}}
	fmt.Println(numIslands(arr))
}

var moves = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
var rowNum int
var colNum int

func numIslands(grid [][]byte) int {
	rowNum = len(grid)
	colNum = len(grid[0])
	count := 0

	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, i, j int) {
	for _, move := range moves {
		newI := i + move[0]
		newJ := j + move[1]
		if newI >= 0 && newJ >= 0 && newI < rowNum && newJ < colNum && grid[newI][newJ] == '1' {
			grid[newI][newJ] = '0'
			dfs(grid, newI, newJ)
		}
	}
}
