package main

import "fmt"

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-1)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Top()
	minStack.Pop()
	// return 0
	fmt.Println(minStack.GetMin()) // return -2
}

type MinStack struct {
	in  [][]int
	min int
}

func Constructor() MinStack {
	return MinStack{
		in: [][]int{},
	}
}

func (this *MinStack) Push(val int) {
	el := []int{val}
	el = append(el, this.min)
	if val < this.min || len(this.in) == 0 {
		this.min = val
	}

	this.in = append(this.in, el)
}

func (this *MinStack) Pop() {
	this.min = this.in[len(this.in)-1][1]
	this.in = this.in[:len(this.in)-1]
}

func (this *MinStack) Top() int {
	return this.in[len(this.in)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
