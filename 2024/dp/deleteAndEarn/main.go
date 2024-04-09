package main

import (
	"fmt"
	"math"
)

func main() {
	resp := deleteAndEarn([]int{2, 2, 3, 3, 3, 4})
	fmt.Println(resp)
}

func deleteAndEarn(nums []int) int {
	return dp(nums)
}

func dp(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 0 {
		return 0
	}

	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	tmpResps := make([]int, 0, 0)
	for k, v := range m {
		var newArr []int
		for i := range arr {
			if arr[i] != k && arr[i] != k-1 && arr[i] != k+1 {
				newArr = append(newArr, arr[i])
			}
		}

		tmpResps = append(tmpResps, v*k+dp(newArr))
	}

	return max(tmpResps)
}

func max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxim := math.MinInt
	for i := range arr {
		if arr[i] > maxim {
			maxim = arr[i]
		}
	}
	return maxim
}
