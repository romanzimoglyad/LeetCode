package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}

type MyQueue struct {
	Stack1 Stack
	Stack2 Stack
}

func Constructor() MyQueue {
	return MyQueue{
		Stack1: Stack{},
		Stack2: Stack{},
	}
}

func (this *MyQueue) Push(x int) {
	for !this.Stack2.IsEmpty() {
		this.Stack1.Push(this.Stack2.Top())
		this.Stack2.Pop()
	}
	this.Stack1.Push(x)
	for !this.Stack2.IsEmpty() {
		this.Stack2.Pop()
	}
	for !this.Stack1.IsEmpty() {
		this.Stack2.Push(this.Stack1.Top())
		this.Stack1.Pop()
	}
}

func (this *MyQueue) Pop() int {

	el := this.Stack2.Top()
	this.Stack2.Pop()

	return el
}

func (this *MyQueue) Peek() int {
	return this.Stack2.Top()

}

func (this *MyQueue) Empty() bool {
	return this.Stack2.IsEmpty()
}

type Stack struct {
	arr []int
}

func ConstructorStack() Stack {
	return Stack{}
}

func (this *Stack) Push(val int) {
	this.arr = append(this.arr, val)

}
func (this *Stack) IsEmpty() bool {
	return len(this.arr) == 0

}
func (this *Stack) Pop() {
	this.arr = this.arr[0 : len(this.arr)-1]

}

func (this *Stack) Top() int {
	return this.arr[len(this.arr)-1]
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
