package main

import "fmt"

func main() {

	obj := Constructor(3)
	_ = obj.EnQueue(1)
	_ = obj.DeQueue()
	_ = obj.Front()
	_ = obj.DeQueue()
	_ = obj.Front()
	_ = obj.Rear()
	_ = obj.EnQueue(1)
	res := obj.IsFull()
	fmt.Println(res)

}

type MyCircularQueue struct {
	size   int
	arr    []int
	curLen int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		size:   k,
		curLen: 0,

		arr: make([]int, 0, k),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {

	if this.size == this.curLen {
		return false
	}
	this.curLen++
	this.arr = append(this.arr, value)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.curLen == 0 {
		return false
	}
	this.curLen--
	this.arr = this.arr[1:]
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.curLen == 0 {
		return -1
	}

	return this.arr[0]
}

func (this *MyCircularQueue) Rear() int {
	if this.curLen == 0 {
		return -1
	}

	return this.arr[len(this.arr)-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.curLen == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.curLen == this.size
}
