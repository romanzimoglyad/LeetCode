package main

import "fmt"

func main() {
	arr := []int{9, 9, 9, 9}
	res := plusOne(arr)
	fmt.Println(res)
}

func plusOne(digits []int) []int {
	l := len(digits)

	for j := l - 1; j >= 0; j-- {
		if digits[j] < 9 {
			digits[j]++
			return digits
		} else {
			digits[j] = 0
		}
	}
	if digits[0] == 0 {
		res := []int{1}
		res = append(res, digits...)
		return res
	}

	return digits
}
