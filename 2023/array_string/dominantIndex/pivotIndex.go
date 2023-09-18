package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	res := dominantIndex(arr)
	fmt.Println(res)
}

func dominantIndex(nums []int) int {
	max := 0
	secondMax := 0
	maxInd := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			secondMax = max
			max = nums[i]
			maxInd = i
		} else {
			if nums[i] > secondMax {
				secondMax = nums[i]
			}
		}
	}
	if secondMax*2 <= max {
		return maxInd
	}
	return -1
}
