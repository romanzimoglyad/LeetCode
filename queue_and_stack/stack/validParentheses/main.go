package main

import "fmt"

func main() {

	fmt.Println(isValid("[][][]{}"))
}

func isValid(s string) bool {
	m := make(map[string]string)
	m["}"] = "{"
	m[")"] = "("
	m["]"] = "["
	if len(s) <= 1 {
		return false
	}
	stack := Constructor()
	r := []rune(s)

	for i := range r {
		el := r[i]
		if !stack.IsEmpty() {
			prevEl := stack.Top()
			if m[string(el)] == string(prevEl) {
				stack.Pop()
			} else {
				stack.Push(el)
			}
		} else {
			stack.Push(el)
		}
	}
	return stack.IsEmpty()
}

type Stack struct {
	next *Stack
	val  rune
}

func Constructor() Stack {
	return Stack{}
}

func (this *Stack) Push(val rune) {
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

func (this *Stack) Top() rune {
	return this.next.val
}
