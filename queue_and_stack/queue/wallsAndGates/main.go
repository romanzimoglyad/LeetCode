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

func wallsAndGates(rooms [][]int) {
	rowSize := len(rooms)
	columnSize := len(rooms[0])

	for row := range rooms {
		for column := range rooms[row] {
			m := make(map[node]struct{})
			queue := newQueue()

			elem := rooms[row][column]
			if elem != wall && elem != gate {
				node := node{
					val:    elem,
					row:    row,
					column: column,
				}
				queue.add(node)
				m[node] = struct{}{}
			}
			step := 0
			for queue.len() > 0 {
				done := false
				size := queue.len()
				for t := 0; t < size; t++ {
					elem := queue.pop()
					if elem.val == gate {
						rooms[row][column] = step
						done = true
						break
					}
					if elem.row+1 < rowSize && rooms[elem.row+1][elem.column] != wall {

						node := node{
							val:    rooms[elem.row+1][elem.column],
							row:    elem.row + 1,
							column: elem.column,
						}
						if _, ok := m[node]; !ok {
							queue.add(node)
							m[node] = struct{}{}
						}

					}
					if elem.row-1 >= 0 && rooms[elem.row-1][elem.column] != wall {

						node := node{
							val:    rooms[elem.row-1][elem.column],
							row:    elem.row - 1,
							column: elem.column,
						}
						if _, ok := m[node]; !ok {
							queue.add(node)
							m[node] = struct{}{}
						}
					}
					if elem.column+1 < columnSize && rooms[elem.row][elem.column+1] != wall {

						node := node{
							val:    rooms[elem.row][elem.column+1],
							row:    elem.row,
							column: elem.column + 1,
						}
						if _, ok := m[node]; !ok {
							queue.add(node)
							m[node] = struct{}{}
						}
					}
					if elem.column-1 >= 0 && rooms[elem.row][elem.column-1] != wall {

						node := node{
							val:    rooms[elem.row][elem.column-1],
							row:    elem.row,
							column: elem.column - 1,
						}
						if _, ok := m[node]; !ok {
							queue.add(node)
							m[node] = struct{}{}
						}
					}
				}
				if done {
					break
				}
				step++
			}

		}

	}

}

type stack struct {
	nodes  []node
	length int
}
type node struct {
	val    int
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
