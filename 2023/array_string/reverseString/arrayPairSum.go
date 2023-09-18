package main

import (
	"fmt"
)

func main() {
	str := []int{2, 7, 11, 15}

	res := twoSum(str, 22)
	fmt.Println(res)
}
func twoSum(numbers []int, target int) []int {

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j]+numbers[i] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{}
}
