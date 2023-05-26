package main

import "fmt"

func main() {
	nums1 := []int{1}

	removeDuplicates(nums1)
	fmt.Println(nums1)
}

func removeDuplicates(nums []int) int {
	j := 1
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[j] = nums[i]
			prev = nums[j]
			j++
		}
	}
	return j
}
