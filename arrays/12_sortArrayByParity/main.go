package main

import "fmt"

func main() {
	nums1 := []int{8, 3, 2, 4, 1, 1, 1}

	n := sortArrayByParity(nums1)
	fmt.Println(n)
	fmt.Println(nums1)
}

/*func sortArrayByParity1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		k := 0
		for nums[i]%2 != 0 && k < len(nums) {
			if nums[i]%2 != 0 {
				last := nums[i]
				for j := i; j < len(nums)-1; j++ {
					nums[j] = nums[j+1]
				}
				nums[len(nums)-1] = last
			}
			k++
		}
	}
	return nums
}*/
func sortArrayByParity(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		k := 1
		for nums[i]%2 != 0 && k < len(nums)-i {
			tmp := nums[i]
			nums[i] = nums[i+k]
			nums[i+k] = tmp
			k++
		}
	}
	return nums
}
