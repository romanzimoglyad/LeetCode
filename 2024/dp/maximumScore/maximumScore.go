package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3}
	multipliers1 := []int{3, 2, 1}

	res := maximumScore(nums1, multipliers1)
	fmt.Println(res)
	nums1 = []int{1}
	multipliers1 = []int{3}
	res = maximumScore(nums1, multipliers1)
	fmt.Println(res)
}

var m map[[2]int]int

func maximumScore(nums []int, multipliers []int) int {
	m = make(map[[2]int]int)
	return dp(0, 0, nums, multipliers)

}

func dp(i, left int, nums, multipliers []int) int {
	n := len(nums)
	right := n - 1 - (i - left)
	if i == len(multipliers) {
		return 0
	}
	if _, ok := m[[2]int{i, left}]; !ok {
		m[[2]int{i, left}] = max(multipliers[i]*nums[left]+dp(i+1, left+1, nums, multipliers), multipliers[i]*nums[right]+dp(i+1, left, nums, multipliers))
	}

	return m[[2]int{i, left}]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
