package main

import "fmt"

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2
}

type MinStack struct {
	ar     []int
	length int
}

func Constructor() MinStack {
	return MinStack{
		ar:     []int{},
		length: 0,
	}
}

func (this *MinStack) Push(val int) {
	this.ar = append(this.ar, val)
	this.length++
}

func (this *MinStack) Pop() {
	this.ar = this.ar[0 : this.length-1]
	this.length--

}

func (this *MinStack) Top() int {
	return this.ar[this.length-1]
}

func (this *MinStack) GetMin() int {
	min := this.ar[0]
	for i := range this.ar {
		if this.ar[i] < min {
			min = this.ar[i]
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
