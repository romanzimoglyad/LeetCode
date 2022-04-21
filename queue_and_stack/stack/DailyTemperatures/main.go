package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{7, 5, 4, 2, 1, 66, 666}))
}

func dailyTemperatures(temperatures []int) []int {
	s := Constructor()
	res := make([]int, len(temperatures))

	for key, value := range temperatures {

		for !s.IsEmpty() && temperatures[s.Top()] < value {
			v := s.Top()
			s.Pop()

			res[v] = key - v
		}
		s.Push(key)

	}

	return res
}

type Stack struct {
	next *Stack
	val  int
}

func Constructor() Stack {
	return Stack{}
}

func (this *Stack) Push(val int) {
	if this.next == nil {
		this.next = &Stack{
			next: nil,
			val:  val,
		}
	} else {

		new := Stack{
			next: this.next,
			val:  val,
		}
		this.next = &new

	}
}
func (this *Stack) IsEmpty() bool {
	return this.next == nil

}
func (this *Stack) Pop() {
	this.next = this.next.next

}

func (this *Stack) Top() int {
	return this.next.val
}
