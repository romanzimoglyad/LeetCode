package main

import "fmt"

func main() {
	testArr := []int{-5, -3, 0, 2}
	res := sortedSquares(testArr)
	fmt.Println(res)
}

func sortedSquares(nums []int) []int {

	l := len(nums)
	resp := make([]int, l)

	leftPos := 0
	rightPos := l - 1
	i := l - 1
	for leftPos <= rightPos {
		leftNew := nums[leftPos] * nums[leftPos]
		rightNew := nums[rightPos] * nums[rightPos]
		if leftNew >= rightNew {
			resp[i] = leftNew
			leftPos++
		} else {
			resp[i] = rightNew
			rightPos--
		}
		i--
	}
	return resp
}
