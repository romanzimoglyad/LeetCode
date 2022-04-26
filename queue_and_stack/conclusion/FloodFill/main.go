package main

import "fmt"

func main() {
	in := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	r := floodFill(in, 1, 1, 2)
	fmt.Println(r)
}

var dirs = [][]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	basicVal := image[sr][sc]
	q := ConstructorQ()
	rl := len(image)
	cl := len(image[0])
	m := make(map[[2]int]struct{})
	m[[2]int{sr, sc}] = struct{}{}
	q.Push([]int{sr, sc})
	for !q.IsEmpty() {
		cur := q.Top()
		image[cur[0]][cur[1]] = newColor
		q.Pop()
		for _, d := range dirs {
			newR := cur[0] + d[0]
			newC := cur[1] + d[1]
			if _, ok := m[[2]int{newR, newC}]; ok || newR < 0 || newC < 0 || newR >= rl || newC >= cl || image[newR][newC] != basicVal {
				continue
			}

			q.Push([]int{newR, newC})
			m[[2]int{newR, newC}] = struct{}{}
		}
	}
	return image
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
