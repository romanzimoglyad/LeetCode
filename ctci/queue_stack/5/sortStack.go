package main

import "fmt"

func main() {
	s := Stack{arr: []int{1, 5, 2, 6, 3, 7, 8, 4}}
	res := sortStack(s)
	fmt.Println(res)
}

type Stack struct {
	arr []int
}

func (s *Stack) Push(i int) {
	s.arr = append(s.arr, i)
}
func (s *Stack) peek() int {
	return s.arr[len(s.arr)-1]

}
func (s *Stack) pop() int {
	el := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return el

}
func (s *Stack) Len() int { return len(s.arr) }

// 12345
func sortStack(stack Stack) Stack {
	finalStack := Stack{}
	for stack.Len() != 0 {
		el := stack.pop()
		if finalStack.Len() == 0 || finalStack.peek() > el {
			finalStack.Push(el)
		} else {
			for finalStack.Len() != 0 && finalStack.peek() < el {
				stack.Push(finalStack.pop())
			}
			finalStack.Push(el)
		}
	}
	return finalStack
}
