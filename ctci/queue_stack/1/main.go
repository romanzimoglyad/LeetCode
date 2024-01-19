package main

import (
	"errors"
	"fmt"
)

func main() {
	s := CreateThreeStuck(3, 1)
	s.push(1, 1)
	s.push(2, 2)
	s.push(3, 3)
	s.pop(3)
	err := s.push(3, 4)
	fmt.Println(err)
	s.peek(1)
	s.peek(2)
	s.peek(3)
	fmt.Println()
}

type ThreeStuck struct {
	arr    []int
	length int

	stackElements []int
}

func CreateThreeStuck(size int, length int) *ThreeStuck {
	arr := make([]int, size*length)
	return &ThreeStuck{
		arr:           arr,
		length:        length,
		stackElements: make([]int, size),
	}
}

func (t *ThreeStuck) push(i int, val int) error {
	if t.stackElements[(i-1)] == t.length {
		return errors.New(fmt.Sprintf("Stack %v is full", i))
	}
	t.arr[(i-1)*t.length+t.stackElements[(i-1)]] = val
	t.stackElements[(i-1)]++
	return nil
}

func (t *ThreeStuck) pop(i int) error {
	if t.stackElements[(i-1)] == 0 {
		return errors.New(fmt.Sprintf("Stack %v is empty", i))
	}
	t.stackElements[(i-1)]--
	return nil
}

func (t *ThreeStuck) peek(i int) (int, error) {
	if t.stackElements[(i-1)] == 0 {
		return 0, errors.New(fmt.Sprintf("Stack %v is empty", i))
	}
	fmt.Printf("Peek %v : %v\n", i, t.arr[(i-1)*t.length+t.stackElements[(i-1)]-1])
	return t.arr[t.stackElements[(i-1)]-1], nil
}
