package main

import "fmt"

func main() {
	q := Queue{stack1: Stack{}, stack2: Stack{}}
	q.push(1)
	q.push(2)
	q.push(3)
	fmt.Println(q.pop())
	q.push(4)
	fmt.Println(q.pop())
	q.push(5)
	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())

}

//12345

type Queue struct {
	stack1 Stack
	stack2 Stack
}

func (q *Queue) push(el int) {
	q.stack1.Push(el)
}

func (q *Queue) pop() int {
	if q.stack2.Len() == 0 {
		for q.stack1.Len() > 0 {
			el := q.stack1.pop()
			q.stack2.Push(el)
		}
	}
	return q.stack2.pop()
}

type Stack struct {
	arr []int
}

func (s *Stack) Push(i int) {
	s.arr = append(s.arr, i)
}

func (s *Stack) pop() int {
	el := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return el

}
func (s *Stack) Len() int { return len(s.arr) }
