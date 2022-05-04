package main

import "fmt"

func main() {
	in := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(in))
}

func maxProfit(prices []int) int {
	sum := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			sum += prices[i+1] - prices[i]
		}
	}
	return sum
}
