package main

import (
	"fmt"
	"math"
)

func main() {
	minStack := Constructor()
	minStack.Push(2)
	minStack.Push(0)
	minStack.Push(3)
	minStack.Push(0)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.GetMin()) // return -2
}

type MinStack struct {
	arr    []int
	length int
	min    []int
}

func Constructor() MinStack {
	return MinStack{arr: make([]int, 0), min: []int{math.MaxInt}}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)
	if val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
	this.length++
}

func (this *MinStack) Pop() {
	el := this.arr[this.length-1]
	if el == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.arr = this.arr[:this.length-1]
	this.length--
}

func (this *MinStack) Top() int {
	return this.arr[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
