package main

import "fmt"

func main() {
	rooms := [][]byte{}
	rooms = [][]byte{
		{earth, earth, earth, earth, water},
		{earth, earth, water, earth, water},
		{earth, earth, water, water, water},
		{water, water, water, water, water}}

	for row := range rooms {
		fmt.Println("")
		for column := range rooms[row] {
			fmt.Print(rooms[row][column])
			fmt.Print(" ")
		}
	}
	fmt.Println("")
	fmt.Println(numIslands(rooms))
}

const (
	earth byte = '1'
	water byte = '0'
)

var directions = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func numIslands(grid [][]byte) int {
	counter := 0
	rowSize := len(grid)
	columnSize := len(grid[0])
	m := make(map[node]struct{})
	for row := range grid {
		for column := range grid[row] {

			queue := newQueue()
			elem := grid[row][column]
			if elem != water {
				node := node{
					row: row,
				}
				if _, ok := m[node]; !ok {
					queue.add(node)
					m[node] = struct{}{}
				} else {
					continue
				}

			} else {
				continue
			}
			for queue.len() > 0 {

				size := queue.len()
				for t := 0; t < size; t++ {
					elem := queue.pop()
					for _, dir := range directions {
						newRow := elem.row + dir[0]
						newColumn := elem.column + dir[1]
						node := node{
							row:    newRow,
							column: newColumn,
						}
						if _, ok := m[node]; ok || newRow < 0 || newColumn < 0 || newRow >= rowSize || newColumn >= columnSize || grid[newRow][newColumn] != earth {
							continue
						}

						m[node] = struct{}{}
						m[node] = struct{}{}
						queue.add(node)
					}
				}
			}
			counter++
		}
	}
	return counter
}

type stack struct {
	nodes  []node
	length int
}
type node struct {
	row    int
	column int
}

func newQueue() stack {
	return stack{
		nodes:  []node{},
		length: 0,
	}
}
func (q *stack) len() int {
	return q.length

}
func (q *stack) add(elem node) {
	q.nodes = append(q.nodes, elem)
	q.length++
}
func (q *stack) pop() node {
	if q.length == 0 {
		return node{}
	}
	res := q.nodes[0]
	q.nodes = q.nodes[1:len(q.nodes)]
	q.length--
	return res
}
