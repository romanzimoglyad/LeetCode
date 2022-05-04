package main

import "fmt"

func main() {
	in := []int{2, 2, 3, 1, 1}

	fmt.Println(singleNumber(in))
}

func singleNumber(nums []int) int {
	var a int = 0
	for _, elem := range nums {
		a ^= elem
	}
	return a
}
