package main

import "fmt"

func main() {
	nums := []int{1}
	t := findTargetSumWays(nums, 1)
	fmt.Println(t)
}

func findTargetSumWays(nums []int, target int) int {

	return dfs(0, nums, target)
}

func dfs(sum int, nums []int, target int) int {
	if sum == target && len(nums) == 0 {
		return 1
	} else if len(nums) == 0 {
		return 0
	}
	p := nums[0]
	nums = nums[1:]

	return dfs(sum+p, nums, target) + dfs(sum-p, nums, target)
}
