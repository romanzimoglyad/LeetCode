package main

import (
	"fmt"
	"math"
)

func main() {
	in := []int{1}
	fmt.Println(removeDuplicates(in))
}

func removeDuplicates(nums []int) int {
	prev := math.MaxInt
	var j int
	for _, elem := range nums {
		if elem != prev {
			prev = elem
			nums[j] = elem
			j++
		}
	}
	return j
}
