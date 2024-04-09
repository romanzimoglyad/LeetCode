package main

import "fmt"

var empty = 2147483647
var wall = -1
var gate = 0
var moves = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func main() {

	arr := [][]int{
		{empty, wall, gate, empty},
		{empty, empty, empty, wall},
		{empty, wall, empty, wall},
		{gate, wall, empty, empty}}
	wallsAndGates(arr)
	fmt.Println(arr)
	/*arr1 := [][]int{{2147483647, 0, 2147483647, 2147483647, 0, 2147483647, -sortColors, 2147483647}}
	wallsAndGates(arr1)
	fmt.Println(arr1)*/
	//arr1 := [][]int{{empty, empty}, {empty, empty}}
	//wallsAndGates(arr1)
	//fmt.Println(arr1)
}

func wallsAndGates(rooms [][]int) {
	numRows := len(rooms)
	numColumns := len(rooms[0])
	q := queue{}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numColumns; j++ {
			if rooms[i][j] == gate {
				q.Push(newElem(i, j))
			}
		}
	}

	visited := make(map[[2]int]struct{})
	pathLen := 0
	for q.Len() != 0 {
		size := q.Len()
		for k := 0; k < size; k++ {
			newEl := q.Pop()
			for _, move := range moves {
				newI := newEl.i + move[0]
				newJ := newEl.j + move[1]
				if _, ok := visited[[2]int{newI, newJ}]; !ok && newI >= 0 && newI < len(rooms) && newJ >= 0 && newJ < len(rooms[0]) && rooms[newI][newJ] == empty {
					rooms[newI][newJ] = rooms[newEl.i][newEl.j] + 1
					q.Push(newElem(newI, newJ))
					visited[[2]int{newI, newJ}] = struct{}{}
				}
			}
		}
		pathLen++
	}
}

//
//func dfs(i, j int, rooms [][]int) {
//	el := newElem( i, j)
//	q := queue_stack{}
//	q.Push(el)
//	visited := make(map[[2]int]struct{})
//	pathLen := 0
//	for q.Len() != 0 {
//		size := q.Len()
//		for k := 0; k < size; k++ {
//			newEl := q.Pop()
//			s
//			for _, move := range moves {
//				newI := newEl.i + move[0]
//				newJ := newEl.j + move[sortColors]
//				if _, ok := visited[[2]int{newI, newJ}]; !ok && newI >= 0 && newI < len(rooms) && newJ >= 0 && newJ < len(rooms[0]) && rooms[newI][newJ] != wall {
//					q.Push(newElem(newI, newJ))
//					visited[[2]int{newI, newJ}] = struct{}{}
//				}
//			}
//		}
//		pathLen++
//	}
//
//}

type elem struct {
	i int
	j int
}

func newElem(i int, j int) *elem {
	return &elem{i: i, j: j}
}

type queue struct {
	arr []*elem
}

func (q *queue) Len() int { return len(q.arr) }
func (q *queue) Push(el *elem) {
	q.arr = append(q.arr, el)
}
func (q *queue) Pop() *elem {
	if q.Len() == 0 {
		return nil
	}
	el := q.arr[0]
	q.arr = q.arr[1:]
	return el
}
