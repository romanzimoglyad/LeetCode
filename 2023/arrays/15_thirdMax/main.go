package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{1, 2, -2147483648}

	res := thirdMax(nums1)
	fmt.Println(res)
}
func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	m1 := math.MinInt
	m2 := math.MinInt
	m3 := math.MinInt
	heat := false
	if len(nums) == 2 {
		return max1(nums[0], nums[1])
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= m1 {
			m1 = nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= m2 && nums[i] != m1 {
			m2 = nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= m3 && nums[i] != m1 && nums[i] != m2 {
			heat = true
			m3 = nums[i]
		}
	}
	if heat {
		return min1(min1(m1, m2), m3)
	} else {
		return max1(m1, m2)
	}
}
func min1(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func max1(i, j int) int {
	if i > j {
		return i
	}
	return j
}
