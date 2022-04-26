package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Top())
}

type MyStack struct {
	q1 Queue
	q2 Queue
}

func Constructor() MyStack {
	return MyStack{
		q1: Queue{},
		q2: Queue{},
	}
}

func (this *MyStack) Push(x int) {
	if this.q1.IsEmpty() {
		this.q1.Push(x)
	} else {
		this.q2.Push(x)
		for !this.q1.IsEmpty() {
			this.q2.Push(this.q1.Top())
			this.q1.Pop()
		}
		this.q2, this.q1 = this.q1, this.q2
	}

}

func (this *MyStack) Pop() int {
	if !this.q1.IsEmpty() {
		el := this.q1.Top()
		this.q1.Pop()
		return el
	} else {
		return 0
	}
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	return this.q1.Top()
}

func (this *MyStack) Empty() bool {
	return this.q1.IsEmpty()
}

type Queue struct {
	arr  []int
	size int
}

func ConstructorStack() Queue {
	return Queue{}
}

func (this *Queue) Push(val int) {
	this.arr = append(this.arr, val)
	this.size++

}
func (this *Queue) IsEmpty() bool {
	return len(this.arr) == 0

}
func (this *Queue) Pop() {
	this.arr = this.arr[1:]
	this.size--

}

func (this *Queue) Size() int {

	return this.size

}
func (this *Queue) Top() int {
	return this.arr[0]
}
