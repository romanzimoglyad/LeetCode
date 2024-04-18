package main

import (
	"fmt"
	"math"
)

func main() {
	r := coinChange([]int{2}, 3)
	fmt.Println(r)
}

func coinChange(coins []int, amount int) int {
	var dp func(left int) int
	memo := make(map[int]int)
	dp = func(left int) int {
		if left == 0 {
			return 0
		}
		if left < 0 {
			return -1
		}

		best := math.MaxInt

		if _, ok := memo[left]; !ok {
			for i := range coins {
				res := dp(left - coins[i])
				if res != -1 {
					best = min(best, res+1)
				}
			}
			if best == math.MaxInt {
				memo[left] = -1
			} else {
				memo[left] = best
			}

		}

		return memo[left]
	}

	return dp(amount)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
