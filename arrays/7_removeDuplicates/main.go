package main

import "fmt"

func main() {
	nums1 := []int{1, 2}

	n := removeDuplicates(nums1)
	fmt.Println(n)
	fmt.Println(nums1)
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
