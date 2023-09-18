package main

import "fmt"

func main() {
	arr := []int{2, 1, -1}
	res := pivotIndex(arr)
	fmt.Println(res)
}

func pivotIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if sum(nums[:i]) == sum(nums[i+1:]) {
			return i
		}
	}
	return -1
}

func sum(in []int) int {
	sum := 0
	for i := range in {
		sum += in[i]
	}
	return sum
}
