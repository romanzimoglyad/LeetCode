package main

import "fmt"

func main() {
	testArr := []int{1, 1, 1, 0, 0, 1, 1, 1, 1}
	res := findMaxConsecutiveOnes(testArr)
	fmt.Println(res)
}

func findMaxConsecutiveOnes(nums []int) int {
	const one = 1
	const zeroLength = 0
	maxOnesInARow := 0
	curNumsInARow := 0
	for _, elem := range nums {
		if elem == one {
			curNumsInARow++
			if curNumsInARow > maxOnesInARow {
				maxOnesInARow = curNumsInARow
			}
		} else {
			curNumsInARow = 0
		}
	}
	return maxOnesInARow
}
