package main

import (
	"fmt"
	"math"
)

func main() {
	in := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	r := updateMatrix(in)
	fmt.Println(r)
}

var dirs = [][]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

func updateMatrix(mat [][]int) [][]int {
	rows := len(mat)
	cols := len(mat[0])
	q := ConstructorQ()
	dist := make([][]int, rows)
	for i := range mat {
		dd := make([]int, cols)
		for k := range mat[0] {
			dd[k] = math.MaxInt
		}
		dist[i] = dd
	}
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				dist[i][j] = 0
				q.Push([]int{i, j})
			}
		}
	}

	for !q.IsEmpty() {
		el := q.Top()
		q.Pop()
		for _, elem := range dirs {
			newR := el[0] + elem[0]
			newC := el[1] + elem[1]
			if newR >= 0 && newC >= 0 && newR < rows && newC < cols {
				if dist[newR][newC] > dist[el[0]][el[1]]+1 {
					dist[newR][newC] = dist[el[0]][el[1]] + 1
					q.Push([]int{newR, newC})
				}
			}
		}
	}
	return dist
}

type Queue struct {
	arr  [][]int
	size int
}

func ConstructorQ() Queue {
	return Queue{}
}

func (this *Queue) Push(val []int) {
	this.arr = append(this.arr, val)
	this.size++

}
func (this *Queue) IsEmpty() bool {
	return len(this.arr) == 0

}
func (this *Queue) Pop() {
	this.arr = this.arr[1:]
	this.size--

}

func (this *Queue) Size() int {

	return this.size

}
func (this *Queue) Top() []int {
	return this.arr[0]
}
