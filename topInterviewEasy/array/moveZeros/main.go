package main

import "fmt"

func main() {
	t := []int{0, 1, 0, 3, 12}
	moveZeroes(t)
	fmt.Println(t)
}

func moveZeroes(nums []int) {
	var j = 0
	for _, elem := range nums {
		if elem != 0 {
			nums[j] = elem
			j++
		} else {
		}
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}
