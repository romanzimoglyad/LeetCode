package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func evalRPN(tokens []string) int {
	stack := Constructor()
	for _, token := range tokens {
		if val, err := strconv.Atoi(token); err == nil {
			stack.Push(val)
		} else {
			switch token {
			case "+":
				val1, val2 := getTwoTopVals(&stack)
				stack.Push(val1 + val2)
			case "-":
				val1, val2 := getTwoTopVals(&stack)
				stack.Push(val2 - val1)
			case "*":
				val1, val2 := getTwoTopVals(&stack)
				stack.Push(val1 * val2)
			case "/":
				val1, val2 := getTwoTopVals(&stack)
				stack.Push(val2 / val1)
			}
		}
	}
	return stack.Top()
}

func getTwoTopVals(stack *stack) (int, int) {
	val1 := stack.Top()
	stack.Pop()
	val2 := stack.Top()
	stack.Pop()
	return val1, val2
}

type stack struct {
	arr    []int
	length int
}

func Constructor() stack {
	return stack{arr: make([]int, 0)}
}

func (this *stack) Push(val int) {
	this.arr = append(this.arr, val)

	this.length++
}

func (this *stack) Pop() {

	this.arr = this.arr[:this.length-1]
	this.length--
}

func (this *stack) Top() int {
	return this.arr[this.length-1]
}

func (this *stack) Len() int {
	return this.length
}
