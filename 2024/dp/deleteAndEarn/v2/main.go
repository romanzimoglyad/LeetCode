package main

import (
	"fmt"
	"math"
)

func main() {
	resp := deleteAndEarn([]int{2, 2, 3, 3, 3, 4})
	fmt.Println(resp)
}

var m = map[int]int{}

var mapNums = map[int]int{}

func deleteAndEarn(nums []int) int {
	m = make(map[int]int)
	mapNums = make(map[int]int)
	max := fillMap(nums)
	return dp(max)
}

func fillMap(arr []int) int {
	max := math.MinInt
	for _, v := range arr {
		mapNums[v] = mapNums[v] + v
		if v > max {
			max = v
		}
	}
	return max
}

func dp(val int) int {
	if val == 0 {
		return 0
	}
	if val == 1 {
		return mapNums[1]
	}
	gain := mapNums[val]

	if _, ok := m[val]; !ok {
		m[val] = max(gain+dp(val-2), dp(val-1))
	}
	return m[val]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
