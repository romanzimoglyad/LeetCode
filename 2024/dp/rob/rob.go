package main

import (
	"fmt"
)

func main() {
	r := rob([]int{2, 7, 9, 3, 1})
	fmt.Println(r)
}

var m = map[int]int{}

func rob(nums []int) int {
	m = make(map[int]int)
	return dp(len(nums)-1, nums)
}

func dp(i int, nums []int) int {
	if i == 0 {
		return nums[0]
	}
	if i == 1 {
		return max(nums[0], nums[i])
	}
	if _, ok := m[i]; !ok {
		m[i] = max(nums[i]+dp(i-2, nums), dp(i-1, nums))
	}

	return m[i]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
