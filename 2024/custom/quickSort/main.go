package main

import (
	"container/list"
	"fmt"
)

func main() {
	test := []int{1, 11, 99, 2, 6, 2}
	fmt.Println(quickSort(test))
	q := &list.List{}
	ptr := q.PushFront("asd")
	q.MoveToFront(ptr)
	l := q.Back()
	fmt.Println(l.Value)
	l = q.Back()
	fmt.Println(l)
}

// 6,4,3,2,1,5
// 4,6,3,2,1,5
// 4,3,2,1,5,6
// 1,5,6 , 1
func partition(a []int, start, end int) ([]int, int) {
	pivot := a[end]
	i := start
	for j := start; j < end; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[end] = a[end], a[i]
	return a, i
}

func qc(a []int, start, end int) []int {
	if start < end {
		a, i := partition(a, start, end)
		a = qc(a, start, i-1)
		a = qc(a, i+1, end)
	}
	return a
}

func quickSort(a []int) []int {
	return qc(a, 0, len(a)-1)
}
