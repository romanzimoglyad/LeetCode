package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	res := spiralOrder(arr)
	fmt.Println(res)
}

func spiralOrder(mat [][]int) []int {
	var res []int
	row := 0
	column := 0
	res = append(res, mat[row][column])
	m := Mover{
		maxRow:    len(mat) - 1,
		maxColumn: len(mat[0]) - 1,
		visited:   make(map[dot]struct{}),
	}
	m.visited[dot{
		row:    row,
		column: column}] = struct{}{}
	for len(m.visited) < len(mat)*len(mat[0]) {
		for {
			if newDot := m.moveRight(dot{
				row:    row,
				column: column}); newDot != nil {
				res = append(res, mat[newDot.row][newDot.column])
				column++
			} else {
				break
			}
		}
		for {
			if newDot := m.moveDown(dot{
				row:    row,
				column: column}); newDot != nil {
				res = append(res, mat[newDot.row][newDot.column])
				row++
			} else {
				break
			}
		}
		for {
			if newDot := m.moveLeft(dot{
				row:    row,
				column: column}); newDot != nil {
				res = append(res, mat[newDot.row][newDot.column])
				column--
			} else {
				break
			}
		}
		for {
			if newDot := m.moveUp(dot{
				row:    row,
				column: column}); newDot != nil {
				res = append(res, mat[newDot.row][newDot.column])
				row--
			} else {
				break
			}
		}
	}
	return res
}

type Mover struct {
	maxRow, maxColumn int
	visited           map[dot]struct{}
}
type dot struct {
	row    int
	column int
}

func (m *Mover) moveRight(in dot) *dot {
	newDot := dot{
		row:    in.row,
		column: in.column + 1,
	}
	if _, ok := m.visited[newDot]; !ok && newDot.column <= m.maxColumn {
		m.visited[newDot] = struct{}{}
		return &newDot
	}
	return nil
}

func (m *Mover) moveDown(in dot) *dot {
	newDot := dot{
		row:    in.row + 1,
		column: in.column,
	}
	if _, ok := m.visited[newDot]; !ok && newDot.row <= m.maxRow {
		m.visited[newDot] = struct{}{}
		return &newDot
	}
	return nil
}
func (m *Mover) moveLeft(in dot) *dot {
	newDot := dot{
		row:    in.row,
		column: in.column - 1,
	}
	if _, ok := m.visited[newDot]; !ok && newDot.column >= 0 {
		m.visited[newDot] = struct{}{}
		return &newDot
	}
	return nil
}
func (m *Mover) moveUp(in dot) *dot {
	newDot := dot{
		row:    in.row - 1,
		column: in.column,
	}
	if _, ok := m.visited[newDot]; !ok && newDot.row >= 0 {
		m.visited[newDot] = struct{}{}
		return &newDot
	}
	return nil
}
