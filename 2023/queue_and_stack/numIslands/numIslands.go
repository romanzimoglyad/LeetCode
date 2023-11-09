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

func numIslands(grid [][]byte) int {
	rowNum := len(grid)
	colNum := len(grid[0])
	count := 0
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			if grid[i][j] == '1' {
				count++
				dfs([2]int{i, j}, grid)
			}
		}
	}
	return count
}

func dfs(el [2]int, grid [][]byte) {
	for _, move := range moves {
		newI := el[0] + move[0]
		newJ := el[1] + move[1]
		if newI >= 0 && newJ >= 0 && newI < len(grid) && newJ < len(grid[0]) && grid[newI][newJ] == '1' {
			grid[newI][newJ] = '0'
			newEl := [2]int{newI, newJ}
			dfs(newEl, grid)
		}
	}
}

type queue struct {
	arr []*[2]int
}

func (q *queue) Len() int { return len(q.arr) }
func (q *queue) Push(el *[2]int) {
	q.arr = append(q.arr, el)
}
func (q *queue) Pop() *[2]int {
	if q.Len() == 0 {
		return nil
	}
	el := q.arr[0]
	q.arr = q.arr[1:]
	return el
}
