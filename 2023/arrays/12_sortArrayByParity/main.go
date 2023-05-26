package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8}

	res := sortArrayByParity(nums1)
	fmt.Println(res)
}
func sortArrayByParity(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			for j := i; j < len(nums); j++ {
				if nums[j]%2 == 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}

	}
	return nums
}
