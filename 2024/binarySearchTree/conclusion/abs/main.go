package main

import (
	"fmt"
	"math"
)

func main() {
	rs := containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3)
	fmt.Println(rs)
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+indexDiff && j < len(nums); j++ {
			if math.Abs(float64(nums[i])-float64(nums[j])) <= float64(valueDiff) {
				return true
			}
		}
	}
	return false
}
