package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 1, 2}

	res := removeDuplicates(arr)
	fmt.Println(res)
}
func removeDuplicates(nums []int) int {
	k := 0
	prev := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if nums[i] != prev {
			nums[k] = nums[i]
			prev = nums[i]
			k++
		}
	}
	return k
}
