package main

import "fmt"

func main() {
	s := "()[]{}{"
	fmt.Println(isValid(s))
}

var links = map[rune]rune{'}': '{', ']': '[', ')': '('} //

func isValid(s string) bool {
	stack := Constructor()
	for _, r := range s {
		if val, ok := links[r]; ok {
			if stack.length == 0 || stack.Top() != val {
				return false
			}
			stack.Pop()
			continue
		}
		stack.Push(r)
	}
	return stack.length == 0
}

type stack struct {
	arr    []rune
	length int
}

func Constructor() stack {
	return stack{arr: make([]rune, 0)}
}

func (this *stack) Push(val rune) {
	this.arr = append(this.arr, val)

	this.length++
}

func (this *stack) Pop() {
	this.arr = this.arr[:this.length-1]
	this.length--
}

func (this *stack) Top() rune {
	return this.arr[this.length-1]
}
