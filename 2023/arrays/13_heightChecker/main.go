package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 4, 3, 4, 5, 6, 7, 8}

	res := heightChecker(nums1)
	fmt.Println(res)
}
func heightChecker(heights []int) int {
	k := 0
	old := make([]int, len(heights))
	_ = copy(old, heights)
	sort.Ints(heights)
	for i := 0; i < len(heights); i++ {
		if heights[i] != old[i] {
			k++
		}
	}
	return k
}
