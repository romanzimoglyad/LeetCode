package main

import (
	"fmt"
)

func main() {
	n := numSquares(13)
	fmt.Println(n)
}
func numSquares(n int) int {
	squaresSlice := []int{}
	for i := 1; i*i <= n; i++ {
		squaresSlice = append(squaresSlice, i*i)
	}

	min := 10000

	var count = 0
	s := newQueue()
	s.add(n)
	s.add(0)
	for s.len() > 1 {
		elem := s.pop()
		if elem == 0 {
			count++
			s.add(0)
			continue
		}
		for _, value := range squaresSlice {
			if elem-value == 0 {
				return count + 1
			}
			if elem-value < 0 {
				break
			}
			if elem-value > 0 {
				s.add(elem - value)
			}
		}
	}

	return min
}

type stack struct {
	nodes  []int
	length int
}

func newQueue() stack {
	return stack{
		nodes:  []int{},
		length: 0,
	}
}
func (q *stack) len() int {
	return q.length

}
func (q *stack) add(elem int) {
	q.nodes = append(q.nodes, elem)
	q.length++
}
func (q *stack) pop() int {
	if q.length == 0 {
		return 0
	}
	res := q.nodes[0]
	q.nodes = q.nodes[1:len(q.nodes)]
	q.length--
	return res
}
