package main

import "fmt"

func main() {
	rooms := [][]int{}
	rooms = [][]int{{empty, wall, gate, empty},
		{empty, empty, empty, wall},
		{empty, wall, empty, wall},
		{gate, wall, empty, empty}}

	wallsAndGates(rooms)
	for row := range rooms {
		fmt.Println("")
		for column := range rooms[row] {
			fmt.Print(rooms[row][column])
			fmt.Print(" ")
		}
	}
	fmt.Println("-------------")
	rooms1 := [][]int{{-1}}
	wallsAndGates(rooms1)
	for row := range rooms1 {
		fmt.Println("")
		for column := range rooms1[row] {
			fmt.Print(rooms1[row][column])
			fmt.Print(" ")
		}
	}
	fmt.Println("-------------")
	rooms2 := [][]int{{empty, empty}, {empty, empty}}
	wallsAndGates(rooms2)
	for row := range rooms2 {
		fmt.Println("")
		for column := range rooms2[row] {
			fmt.Print(rooms2[row][column])
			fmt.Print(" ")
		}
	}
}

const (
	empty = 2147483647
	wall  = -1
	gate  = 0
)

var directions = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func wallsAndGates(rooms [][]int) {
	rowSize := len(rooms)
	columnSize := len(rooms[0])
	queue := newQueue()
	for row := range rooms {
		for column := range rooms[row] {
			if rooms[row][column] == gate {
				queue.add(node{
					row:    row,
					column: column,
				})
			}
		}
	}

	for queue.len() > 0 {
		elem := queue.pop()
		for _, dir := range directions {

			newRow := elem.row + dir[0]
			newColumn := elem.column + dir[1]
			if newRow < 0 || newColumn < 0 || newRow >= rowSize || newColumn >= columnSize || rooms[newRow][newColumn] != empty {
				continue
			}
			rooms[newRow][newColumn] = rooms[elem.row][elem.column] + 1
			queue.add(node{
				row:    newRow,
				column: newColumn})
		}

	}

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
