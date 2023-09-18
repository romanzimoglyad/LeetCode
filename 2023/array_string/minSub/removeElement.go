package main

import (
	"fmt"
)

func main() {
	str := []int{1, 1, 0, 1, 1, 1}

	res := findMaxConsecutiveOnes(str)
	fmt.Println(res)
}
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			cur++
			if cur >= max {
				max = cur
			}
		} else {
			cur = 0
		}
	}
	return max
}
