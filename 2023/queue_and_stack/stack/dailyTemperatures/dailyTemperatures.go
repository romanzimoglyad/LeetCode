package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(temperatures []int) []int {
	s := Constructor()
	resp := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for el := s.Top(); el != nil && el.temp < temperatures[i]; el = s.Top() {
			resp[el.pos] = i - el.pos
			s.Pop()

		}
		s.Push(&node{temp: temperatures[i], pos: i})
	}

	return resp
}

type node struct {
	temp int
	pos  int
}

type stack struct {
	arr    []*node
	length int
}

func Constructor() stack {
	return stack{arr: make([]*node, 0)}
}

func (this *stack) Push(val *node) {
	this.arr = append(this.arr, val)

	this.length++
}

func (this *stack) Pop() {
	if this.length == 0 {
		return
	}
	this.arr = this.arr[:this.length-1]
	this.length--
}

func (this *stack) Top() *node {
	if this.length == 0 {
		return nil
	}
	return this.arr[this.length-1]
}
