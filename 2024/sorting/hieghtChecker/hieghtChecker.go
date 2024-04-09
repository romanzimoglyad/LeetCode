package main

import "fmt"

// 1 1 1 2 3 4

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
}

func heightChecker(heights []int) int {
	hasSwapped := true
	old := make([]int, len(heights))
	copy(old, heights)
	for hasSwapped {
		hasSwapped = false
		for i := 0; i < len(heights)-1; i++ {
			if heights[i+1] < heights[i] {
				heights[i+1], heights[i] = heights[i], heights[i+1]
				hasSwapped = true
			}
		}
	}
	res := 0
	for i := range heights {
		if heights[i] != old[i] {
			res++
		}
	}
	return res
}
