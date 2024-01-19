package main

import "fmt"

func main() {
	myQueue := Constructor()
	myQueue.Push(1)              // queue_stack is: [1]
	myQueue.Push(2)              // queue_stack is: [1, 2] (leftmost is front of the queue_stack)
	fmt.Println(myQueue.Peek())  // return 1
	fmt.Println(myQueue.Pop())   // return 1, queue_stack is [2]
	fmt.Println(myQueue.Empty()) // return false
}

type MyQueue struct {
	stack1      stack
	stack2      stack
	whereToPick int
	length      int
}

func Constructor() MyQueue {
	return MyQueue{
		stack1:      ConstructorStack(),
		stack2:      ConstructorStack(),
		whereToPick: 0,
	}
}

func (this *MyQueue) Push(x int) {
	this.stack1.Push(x)
	this.length++
}

func (this *MyQueue) Pop() int {
	if this.stack2.length == 0 {
		for this.stack1.length > 0 {
			this.stack2.Push(this.stack1.Top())
			this.stack1.Pop()
		}
	}
	resp := this.stack2.Top()
	this.stack2.Pop()
	this.length--
	return resp

}

func (this *MyQueue) Peek() int {
	if this.stack2.length == 0 {
		for this.stack1.length > 0 {
			this.stack2.Push(this.stack1.Top())
			this.stack1.Pop()
		}
	}
	return this.stack2.Top()

}

func (this *MyQueue) Empty() bool {
	return this.length == 0
}

type stack struct {
	arr    []int
	length int
}

func ConstructorStack() stack {
	return stack{arr: make([]int, 0)}
}

func (this *stack) Push(val int) {
	this.arr = append(this.arr, val)

	this.length++
}

func (this *stack) Pop() {
	this.arr = this.arr[:this.length-1]
	this.length--
}

func (this *stack) Top() int {
	return this.arr[this.length-1]
}
