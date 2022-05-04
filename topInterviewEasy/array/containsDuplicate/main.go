package main

import "fmt"

func main() {
	in := []int{1, 2, 1, 4, 5, 6, 7}

	fmt.Println(containsDuplicate(in))
}
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for i := range nums {
		if m[nums[i]] {
			return true
		}
		m[nums[i]] = true
	}
	return false
}
