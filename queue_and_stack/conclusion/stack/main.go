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
	q1  Queue
	q2  Queue
	top int
}

func Constructor() MyStack {
	return MyStack{
		q1:  Queue{},
		q2:  Queue{},
		top: 0,
	}
}

func (this *MyStack) Push(x int) {
	if this.Empty() || !this.q1.IsEmpty() {
		this.q1.Push(x)
	} else {
		this.q1.Push(x)
	}
	this.top = x

}

func (this *MyStack) Pop() int {
	if !this.q1.IsEmpty() {

		for i := this.q1.Size(); i > 1; i = this.q1.Size() {
			this.q2.Push(this.q1.Top())

			if i == 2 {
				this.top = this.q1.Top()
			}
			this.q1.Pop()
		}

		lastEl := this.q1.Top()
		this.q1.Pop()
		return lastEl
	} else {

		for i := this.q2.Size(); i > 1; i = this.q1.Size() {
			this.q1.Push(this.q2.Top())

			if i == 2 {
				this.top = this.q2.Top()
			}
			this.q2.Pop()
		}
		lastEl := this.q2.Top()
		this.q2.Pop()

		return lastEl
	}
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	return this.top
}

func (this *MyStack) Empty() bool {
	return this.q1.IsEmpty() && this.q2.IsEmpty()
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
