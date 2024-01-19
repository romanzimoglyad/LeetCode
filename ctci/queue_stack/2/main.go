package main

import (
	"fmt"
	"math"
)

func main() {
	m := MinStack{
		main:    []int{},
		min:     []int{},
		length:  0,
		lastMin: math.MaxInt,
	}
	m.Push(7)
	m.Push(6)
	m.Push(8)
	m.Push(3)
	fmt.Println(m.Min())
	fmt.Println(m.Pop())
	fmt.Println(m.Min())
	fmt.Println(m.Pop())
	fmt.Println(m.Min())
	fmt.Println(m.Pop())
	fmt.Println(m.Min())
	fmt.Println(m.Pop())
}

type MinStack struct {
	main    []int
	min     []int
	length  int
	lastMin int
}

func (s *MinStack) Push(elem int) {
	if elem < s.lastMin {
		s.min = append(s.min, elem)
		s.lastMin = elem
	}
	s.length++
	s.main = append(s.main, elem)
}
func (s *MinStack) Pop() int {
	val := s.main[s.length-1]
	s.length--
	if s.lastMin == val {

		s.min = s.min[:len(s.min)-1]
		if s.length == 0 {
			s.lastMin = math.MaxInt
		} else {
			s.lastMin = s.min[len(s.min)-1]
		}
	}
	s.main = s.main[:len(s.main)-1]

	return val
}

func (s *MinStack) Peek() int {
	return s.main[s.length-1]
}
func (s *MinStack) Len() int { return len(s.main) }

func (s *MinStack) Min() int {
	return s.lastMin
}
