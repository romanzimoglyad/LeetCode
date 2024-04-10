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

func maximumScore(nums []int, multipliers []int) int {

	//dp := make([][]int, len(multipliers))

	//	for i := len(multipliers) - 1; i >= 0; i-- {
	//
	// for k:= len(multipliers[i]);
	//
	//		max(multipliers[i]*nums[left]+dp(i+1, left+1, nums, multipliers), multipliers[i]*nums[right]+dp(i+1, left, nums, multipliers))
	//	}
	//
	//	return dp[0][0]
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
