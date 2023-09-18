package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 1, 2, 0, 0, 2, 1, 4}

	moveZeroes(arr)
	fmt.Println(arr)
}
func moveZeroes(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}
	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}
