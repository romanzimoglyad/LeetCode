package main

import (
	"fmt"
	"math"
)

func main() {
	r := minDifficulty([]int{1, 1, 1, 6, 6}, 2)
	fmt.Println(r)
}

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	maxRemaining := make([]int, len(jobDifficulty))
	for i := range jobDifficulty {
		maxRemaining[i] = finMaxRemaining(i, jobDifficulty)
	}
	memo := make(map[[2]int]int)

	var dp func(index, day int) int
	dp = func(index, day int) int {
		if day == d {
			return maxRemaining[index]
		}
		best := math.MaxInt
		jobMax := 0
		if _, ok := memo[[2]int{index, day}]; !ok {
			for j := index; j < n-(d-day); j++ {
				if jobDifficulty[j] > jobMax {
					jobMax = jobDifficulty[j]
				}
				best = min(best, jobMax+dp(j+1, day+1))
			}
			memo[[2]int{index, day}] = best
		}

		return memo[[2]int{index, day}]
	}

	return dp(0, 1)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func finMaxRemaining(i int, arr []int) int {
	max := arr[i]
	for j := i; j < len(arr); j++ {
		if arr[j] > max {
			max = arr[j]
		}
	}
	return max
}
