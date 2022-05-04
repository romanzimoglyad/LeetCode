package main

import "fmt"

func main() {
	in := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(in))
}

func maxProfit(prices []int) int {
	var curSum int
	prev := -1
	cur := -1
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] < prices[i] {
			if prev != -1 && cur != -1 {
				curSum += cur - prev
				prev, cur = -1, -1
			}
			continue
		} else {
			if prev == -1 {
				prev = prices[i]
			}
			cur = prices[i+1]
		}
	}
	if prev != -1 && cur != -1 {
		curSum += cur - prev

	}
	return curSum
}
