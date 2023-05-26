package main

import "fmt"

func main() {
	nums1 := []int{1, 3, 3, 2, 3}

	removeElement(nums1, 3)
	fmt.Println(nums1)
}

func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
		} else {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
