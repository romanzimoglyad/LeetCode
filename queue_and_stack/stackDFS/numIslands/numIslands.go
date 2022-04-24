package main

import "fmt"

const (
	earth byte = '1'
	water byte = '0'
)

var dir = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func main() {
	rooms := [][]byte{
		{earth, earth, earth, earth, water},
		{earth, earth, water, earth, water},
		{earth, earth, water, water, water},
		{water, water, water, water, water}}
	t := numIslands(rooms)
	fmt.Println(t)
}

func numIslands(grid [][]byte) int {
	c := 0

	for row := range grid {
		for column := range grid[row] {
			el := grid[row][column]
			if el == earth {
				c++
				dfs(grid, row, column)
			}
		}
	}
	return c
}
func dfs(grid [][]byte, row int, column int) {
	rowL := len(grid)
	columnL := len(grid[0])
	for d := range dir {
		newRow := row + dir[d][0]
		newColumn := column + dir[d][1]
		if newRow < 0 || newColumn < 0 || newRow >= rowL || newColumn >= columnL || grid[newRow][newColumn] != earth {
			continue
		} else {
			grid[newRow][newColumn] = water
			dfs(grid, newRow, newColumn)
		}
	}
}
