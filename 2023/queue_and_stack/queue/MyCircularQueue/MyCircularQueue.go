package main

import "fmt"

func main() {
	q := Constructor(2)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)

	fmt.Println(q.Front())
	fmt.Println(q.Rear())
	q.DeQueue()
	q.DeQueue()
	fmt.Println(q.Front())
	fmt.Println(q.Rear())
	q.EnQueue(6)
	q.EnQueue(7)
	fmt.Println(q.Front())
	fmt.Println(q.Rear())
	q.DeQueue()
	q.DeQueue()
	q.DeQueue()
	fmt.Println(q.Front())
	fmt.Println(q.Rear())
}

type MyCircularQueue struct {
	startInd int
	endInd   int
	arr      []int
	size     int
	cap      int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		startInd: 0,
		endInd:   -1,
		arr:      make([]int, k+1, k+1),
		cap:      k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.size == this.cap {
		return false
	}
	if this.endInd < this.cap-1 {
		this.endInd++
		this.arr[this.endInd] = value
	} else {
		this.endInd = 0
		this.arr[this.endInd] = value
	}
	this.size++
	fmt.Println(this.arr)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.size == 0 {
		return false
	}
	if this.startInd < this.cap-1 {
		this.startInd++
	} else {
		this.startInd = 0
	}
	this.size--
	fmt.Println(this.arr)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.size == 0 {
		return -1
	}
	return this.arr[this.startInd]
}

func (this *MyCircularQueue) Rear() int {
	if this.size == 0 {
		return -1
	}
	return this.arr[this.endInd]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.cap
}
