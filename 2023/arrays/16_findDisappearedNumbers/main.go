package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 4, 5, 6, 7, 9}

	res := findDisappearedNumbers(nums1)
	fmt.Println(res)
}
func findDisappearedNumbers(nums []int) []int {
	var res []int
	m := make(map[int]struct{})
	n := len(nums) + 1
	for i := range nums {
		m[nums[i]] = struct{}{}
	}
	for i := 1; i < n; i++ {
		if _, ok := m[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}
