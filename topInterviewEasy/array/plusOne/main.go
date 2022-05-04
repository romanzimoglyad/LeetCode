package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9}))
}
func plusOne(digits []int) []int {

	for j := len(digits) - 1; j >= 0; j-- {
		if digits[j] == 9 {

			digits[j] = 0
		} else {
			digits[j]++

			return digits
		}
	}
	res := []int{1}

	res = append(res, digits...)
	return res

}
