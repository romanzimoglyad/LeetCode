package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 0, 1, 1, 0}

	res := findMaxConsecutiveOnes(nums1)
	fmt.Println(res)
}
func findMaxConsecutiveOnes(nums []int) int {
	max := 0

	for i := 0; i < len(nums); i++ {
		num0 := 0
		for j := i; j < len(nums); j++ {
			if nums[j] == 0 {
				num0++
			}
			if num0 <= 1 {
				max = max1(max, j-i+1)
			}
		}
	}
	return max
}

func max1(i, j int) int {
	if i > j {
		return i
	}
	return j
}
