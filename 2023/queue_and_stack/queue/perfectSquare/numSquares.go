package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(7168))
}

func numSquares(n int) int {
	squares := initSquares(10000)
	q := queue{}
	min := math.MaxInt
	q.Push(&n)
	count := 0
	for q.Len() > 0 {
		size := q.Len()
		count++
		for i := 0; i < size; i++ {
			el := q.Pop()
			for j := range squares {
				if diff := *el - squares[j]; diff == 0 {
					return count
				} else if diff := *el - squares[j]; diff > 0 {
					q.Push(&diff)
				} else {
					break
				}
			}

		}

	}
	return min
}

func minF(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func initSquares(n int) []int {
	res := make([]int, 0, int(math.Sqrt(float64(n))))
	for i := 1; i*i <= n; i++ {
		res = append(res, i*i)
	}
	return res
}

type queue struct {
	arr []*int
}

func (q *queue) Len() int { return len(q.arr) }
func (q *queue) Push(el *int) {
	q.arr = append(q.arr, el)
}
func (q *queue) Pop() *int {
	if q.Len() == 0 {
		return nil
	}
	el := q.arr[0]
	q.arr = q.arr[1:]
	return el
}
