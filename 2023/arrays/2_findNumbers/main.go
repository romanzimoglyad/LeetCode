package main

import (
	"fmt"
	"strconv"
)

func main() {
	testArr := []int{555, 901, 482, 1771}
	res := findNumbers(testArr)
	fmt.Println(res)
}

func findNumbers(nums []int) int {
	evenDigitsCount := 0
	for i := range nums {
		str := strconv.Itoa(nums[i])
		if len(str)%2 == 0 {
			evenDigitsCount++
		}
	}
	return evenDigitsCount
}
