package main

import "fmt"

func main() {
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}

	n := removeElement(nums1, 2)
	fmt.Println(n)
	fmt.Println(nums1)
}

func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
