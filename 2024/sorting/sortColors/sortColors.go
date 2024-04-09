package main

import "fmt"

func main() {
	nums := []int{0, 2, 1, 2, 1, 0, 0, 1}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {

	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}
