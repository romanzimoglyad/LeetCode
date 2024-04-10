package main

import (
	"fmt"
	"math"
)

func main() {
	resp := deleteAndEarn([]int{3, 4, 2})
	fmt.Println(resp)
}

var mapNums = map[int]int{}

func deleteAndEarn(nums []int) int {

	mapNums = make(map[int]int)

	maximum := fillMap(nums)
	oneBack := mapNums[1]
	twoBack := 0

	for i := 2; i <= maximum; i++ {
		tmp := oneBack
		oneBack = max(mapNums[i]+twoBack, oneBack)
		twoBack = tmp
	}

	return oneBack
}

func fillMap(arr []int) int {
	max := math.MinInt
	for _, v := range arr {
		mapNums[v] = mapNums[v] + v
		if v > max {
			max = v
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
