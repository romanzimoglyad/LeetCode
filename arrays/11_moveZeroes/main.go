package main

import "fmt"

func main() {
	nums1 := []int{0, 1, 0, 0, 0, 3, 12}

	moveZeroes(nums1)

	fmt.Println(nums1)
}

func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
		} else {
			nums[j] = nums[i]
			j++
		}
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}

}
