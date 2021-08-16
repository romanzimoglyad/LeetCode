package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{2, 2, 2, 3, 3, 1}

	n := thirdMax(nums1)
	fmt.Println(n)
	fmt.Println(nums1)
}

func thirdMax(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	k := removeDuplicates(nums)
	if k < 3 {
		return nums[0]
	}
	return nums[2]
}
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 1
	for i := 0; i < len(nums)-1; i++ {
		new := nums[i+1]
		prev := nums[i]
		if new != prev {
			nums[j] = new
			j++
		}
	}
	return j
}
