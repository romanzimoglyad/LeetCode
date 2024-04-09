package main

import "fmt"

func main() {
	r := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	fmt.Println(r)
}

var m = map[int]int{}

func minCostClimbingStairs(cost []int) int {
	m = make(map[int]int)
	cost = append(cost, 0)
	return dp(len(cost)-1, cost)
}

func dp(i int, nums []int) int {
	if i == 0 {
		return nums[0]
	}
	if i == 1 {
		return nums[1]
	}
	if _, ok := m[i]; !ok {
		m[i] = nums[i] + min(dp(i-1, nums), dp(i-2, nums))
	}

	return m[i]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
