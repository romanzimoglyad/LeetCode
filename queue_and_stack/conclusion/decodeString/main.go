package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "100[leetcode]"
	s := decodeString(str)
	fmt.Println(s)
}
func decodeString(s string) string {
	runeStr := []rune(s)
	mDigit := map[string]struct{}{"0": {},
		"1": {},
		"2": {},
		"3": {},
		"4": {},
		"5": {},
		"6": {},
		"7": {},
		"8": {},
		"9": {}}

	stack := ConstructorStack()
	for _, elem := range runeStr {
		if string(elem) != "]" {
			stack.Push(string(elem))
		} else {
			curStr := ""
			s2 := ConstructorStack()
			for stack.Top() != "[" {
				s2.Push(stack.Top())
				stack.Pop()
			}
			for !s2.IsEmpty() {
				curStr += s2.Top()
				s2.Pop()
			}

			stack.Pop()
			curDigit := ""

			s3 := ConstructorStack()
			for !stack.IsEmpty() {
				if _, ok := mDigit[stack.Top()]; ok {
					s3.Push(stack.Top())

					stack.Pop()
				} else {
					break
				}

			}
			for !s3.IsEmpty() {
				curDigit += s3.Top()
				s3.Pop()
			}

			d, err := strconv.Atoi(curDigit)
			if err != nil {
				panic("asd")
			}
			nesStr := ""
			for i := 0; i < d; i++ {
				nesStr += curStr
			}
			stack.Push(nesStr)
		}

	}
	return stack.Print()
}

type Stack struct {
	arr []string
}

func ConstructorStack() Stack {
	return Stack{}
}
func (this *Stack) Print() string {
	return strings.Join(this.arr, "")

}
func (this *Stack) Push(val string) {
	this.arr = append(this.arr, val)

}
func (this *Stack) IsEmpty() bool {
	return len(this.arr) == 0

}
func (this *Stack) Pop() {
	this.arr = this.arr[0 : len(this.arr)-1]

}

func (this *Stack) Top() string {
	return this.arr[len(this.arr)-1]
}
