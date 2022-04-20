package main

import "fmt"

func main() {
	minStack := Constructor()
	minStack.Push(2)
	minStack.Push(0)
	minStack.Push(3)
	fmt.Println(minStack.GetMax()) // return 3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMax()) // return 2
}

type MaxStack struct {
	next *MaxStack
	max  int
	val  int
}

func Constructor() MaxStack {
	return MaxStack{}
}

func (this *MaxStack) Push(val int) {
	if this.next == nil {
		this.next = &MaxStack{
			next: nil,
			max:  val,
			val:  val,
		}
	} else {

		newMax := this.next.max
		if val > this.next.max {
			newMax = val
		}
		new := MaxStack{
			next: this.next,
			max:  newMax,
			val:  val,
		}
		this.next = &new

	}
}

func (this *MaxStack) Pop() {
	this.next = this.next.next

}

func (this *MaxStack) Top() int {
	return this.next.val
}

func (this *MaxStack) GetMax() int {
	return this.next.max
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
