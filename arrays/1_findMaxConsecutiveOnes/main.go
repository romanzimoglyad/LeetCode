package main

import "fmt"

func main() {
	v := []int{1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 2}
	t := findMaxConsecutiveOnes(v)
	fmt.Println(t)
}

func findMaxConsecutiveOnes(nums []int) int {
	tempRes := 0
	biggestRes := 0
	for i := range nums {
		if nums[i] == 1 {
			tempRes++
		} else {
			if tempRes > biggestRes {
				biggestRes = tempRes
			}
			tempRes = 0
		}
	}
	if tempRes > biggestRes {
		biggestRes = tempRes
	}
	return biggestRes
}
