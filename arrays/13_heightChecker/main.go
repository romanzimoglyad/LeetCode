package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{5, 1, 2, 3, 4}

	n := heightChecker(nums1)
	fmt.Println(n)
	fmt.Println(nums1)
}

func heightChecker(heights []int) int {
	in := make([]int, len(heights))
	copy(in, heights)
	sort.Ints(heights)

	i := 0
	j := 0
	for i < len(heights) {
		if in[i] != heights[i] {
			j++
		}
		i++
	}
	return j
}
